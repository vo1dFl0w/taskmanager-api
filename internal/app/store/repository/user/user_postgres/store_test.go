package user_postgres_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	 databaseURL = os.Getenv("DATABASE_URL")
	 if databaseURL == "" {
		databaseURL = "host=localhost dbname=todo-api-db_test sslmode=disable"
	 }

	 os.Exit(m.Run())
}