package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"tutorial.sqlc.dev/app/utils"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: " + err.Error())
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
