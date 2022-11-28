package store

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	// ...
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=admin password=lillol000 dbname=restapi_dev sslmode=disable"
	}
	os.Exit(m.Run())
}
