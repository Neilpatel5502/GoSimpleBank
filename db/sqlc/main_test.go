package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	testDb, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the db:", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
