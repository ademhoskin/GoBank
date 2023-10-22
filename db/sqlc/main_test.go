package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
	_ "github.com/jackc/pgconn"
)

var testQueries *Queries

func TestMain (m *testing.M) {
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}