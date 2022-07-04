package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepo

// setup main test
func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db") // remove existing DB for migration test
	path := "./testdata/sql.db"

	// connect SQLite DB
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepo(db)
	os.Exit(m.Run())
}
