package db

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"cheat/cli/cmd/exceptions"
	"cheat/cli/models"
	"cheat/cli/utils"
)

var database *sql.DB
var tx *sql.Tx

type _Cheat = models.Cheat

// Cleanup : close database connection and commit/rollback
func Cleanup() {
	err := tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			panic(exceptions.CheatException("Could not rollback database transaction", err))
		}
		panic(exceptions.CheatException("Could not commit database transaction", err))
	}
	err = database.Close()
	if err != nil {
		panic(exceptions.CheatException("Could not close database connection", err))
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

	regex := func(re string, s string, ignoreCase bool) (bool, error) {
		caseIgnore := map[bool]string{true: "(?i)", false: ""}[ignoreCase]
		return regexp.MatchString(caseIgnore+re, s)
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
		panic(exceptions.CheatException("Could not establish database connection", err))
	}

	// Force db to make a new connection in pool
	// by putting the original in a transaction
	tx, err = database.Begin()
	if err != nil {
		panic(exceptions.CheatException("Could not start a new transaction in database pool", err))
	}

	// Initialize table
	_, err = database.Exec(`
    CREATE TABLE IF NOT EXISTS cheat (
      name TEXT PRIMARY KEY,
      created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      description TEXT,
      weight INTEGER NOT NULL DEFAULT 0
    )
  `)
	if err != nil {
		panic(exceptions.CheatException("Could not initialize database with table \"cheat\"", err))
	}
}
