package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	connection, error := sql.Open(dbDriver, dbSource)

	if error != nil {
		log.Fatal("cannot connect to database", error)
	}

	testQueries = New(connection)

	os.Exit(m.Run())

}
