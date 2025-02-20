package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Neilpatel5502/GoSimpleBank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadCondig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDb, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the db:", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
