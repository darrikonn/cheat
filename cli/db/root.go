package db

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	"github.com/mattn/go-sqlite3"
)

var database *sql.DB
var tx *sql.Tx

// Cleanup : close database connection and commit/rollback
func Cleanup() {
	err := tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
		panic(err)
	}
	err = database.Close()
	if err != nil {
		panic(err)
	}
}

// Setup : sets up the database and both
// 1. Creates the database
// 2. Initializes the necessary tables
func Setup() {
	var err error

	databasePath := utils.ResolvePath(viper.GetString("database"))
	if !utils.FileExists(databasePath) {
		utils.CreateFile(databasePath)
	}

	sql.Register(
		"sqlite3_with_regexp",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		},
	)

	database, err = sql.Open(
		"sqlite3_with_regexp",
		fmt.Sprintf("file:%s?mode=rw", databasePath),
	)
	if err != nil {
		panic(err)
	}

	// Force db to make a new connection in pool
	// by putting the original in a transaction
	tx, err = database.Begin()
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(`
    CREATE TABLE IF NOT EXISTS cheat (
      id TEXT PRIMARY KEY,
      created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      command TEXT NOT NULL,
      name TEXT,
      description TEXT,
      weight INTEGER NOT NULL DEFAULT 0
    )
  `)
	if err != nil {
		panic(err)
	}
}
