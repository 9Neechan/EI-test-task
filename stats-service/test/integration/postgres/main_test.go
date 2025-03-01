package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	//"github.com/9Neechan/EI-test-task/util"
	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	cfg "github.com/9Neechan/EI-test-task/stats-service/internal/config"
	
	_ "github.com/lib/pq"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := cfg.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
