package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_USR")
	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit((m.Run()))

}
