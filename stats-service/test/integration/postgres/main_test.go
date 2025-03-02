package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	cfg "github.com/9Neechan/EI-test-task/stats-service/internal/config"
	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

// Declaring global variables for test queries and database connection
var testQueries *sqlc.Queries
var testDB *sql.DB

// Function to setup and run tests
func TestMain(m *testing.M) {
	var err error

	// Loading configuration for the test
	config, err := cfg.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Opening a database connection
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Creating a new instance of Queries
	testQueries = sqlc.New(testDB)

	// Running the tests
	os.Exit(m.Run())
}
