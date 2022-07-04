package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepo

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"

	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepo(db)
	os.Exit(m.Run())
}
