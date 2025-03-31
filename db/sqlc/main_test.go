package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/FilledEther20/Reg_Bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot set config file: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	// Ensure database connection is closed after all tests run
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}
