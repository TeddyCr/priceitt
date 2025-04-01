// +build migration

package main

import (
	"os"
	"testing"

	"github.com/TeddyCr/priceitt/service/test"
	"github.com/TeddyCr/priceitt/service/utils/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("EDGE_AUTHORIZATION_SERVER_CONFIG_FILE_PATH", "cmd/testdata/config.yaml")
	psqlContainer := test.DefaultPostgresTestHandler()
	handlers := []test.ITestHandler{psqlContainer}
	test.SetUp(handlers)
	m.Run()
	test.TearDown(handlers)
}

func TestMigrate(t *testing.T) {
	run()
	config := getConfigModel()

	db := database.Connect(config.Database)
	validateMigrations(db, t)
}

func validateMigrations(db *sqlx.DB, t *testing.T) {
	expectedTableNames := []string{"database_migration_logs", "users"}

	query := "SELECT COUNT(*) FROM DATABASE_MIGRATION_LOGS;"
	row := db.QueryRow(query)
	var count int
	err := row.Scan(&count)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, count)

	query = "SELECT table_name FROM information_schema.tables WHERE table_catalog = 'edge_authorization_server' AND table_schema = 'public';"
	rows, _ := db.Query(query)
	tables := []string{}
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			t.Error(err)
		}
		tables = append(tables, tableName)
	}
	assert.ElementsMatch(t, expectedTableNames, tables)
}
