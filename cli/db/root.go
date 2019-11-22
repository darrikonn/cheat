package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/mattn/go-sqlite3"
)

var database *sql.DB
var tx *sql.Tx

func Cleanup() {
	database.Close()
	err := tx.Commit()
	if err != nil {
		tx.Rollback()
	}
}
func init() {
	var err error

	homeDirectory, _ := os.UserHomeDir()
	regex := func(re, s string) (bool, error) {
		return regexp.MatchString(re, s)
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
		fmt.Sprintf("file:%s?mode=rw", homeDirectory+"/.cheetsheet.db"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// database.SetMaxOpenConns(1)
	// database.Exec("PRAGMA journal_mode=WAL;")

	// Force db to make a new connection in pool
	// by putting the original in a transaction
	tx, err = database.Begin()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
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

	// if err != nil {
	//   panic(err)
	// }
}
