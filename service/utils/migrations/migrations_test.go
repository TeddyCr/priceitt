package migrations_test

import (
	"os"
	"testing"

	"github.com/TeddyCr/priceitt/service/models"
	"github.com/TeddyCr/priceitt/service/utils/migrations"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func tearDown() {
	// Remove test db
	os.Remove("testdata/test.db") //nolint:errcheck
}

func TestMain(m *testing.M) {
	m.Run()
	tearDown()
}

// TestExecMigration tests the migration of schema and data files
// We'll create a test db and run the migration on it
// We'll then valiudate the migration by checking the metadata table
// and the data in the user table
func TestExecMigration(t *testing.T) {
	// 1. Run migration
	migrationConfig := models.MigrationConfig{
		SchemaPath:   "/utils/migrations/testdata/schema",
		DataPath:     "/utils/migrations/testdata/data",
		MetadataPath: "/utils/migrations/testdata/metadata",
		MetadataQueries: models.MigrationMetadataQueries{
			InsertMetadataQuery:           "INSERT INTO DATABASE_MIGRATION_LOGS (version, query, checksum, execution_time) VALUES ($1, $2, $3, $4)",
			SelectMigrationVersionsQuery:  "SELECT version FROM DATABASE_MIGRATION_LOGS",
			SelectMigrationChecksumsQuery: "SELECT checksum FROM DATABASE_MIGRATION_LOGS",
		},
		Force:          false,
		CheckIntegrity: false,
	}

	dbConfig := models.DatabaseConfig{
		DriverClass:      "sqlite3",
		ConnectionString: "testdata/test.db",
	}

	migrations.ExecMigration(migrationConfig, dbConfig)

	db, err := sqlx.Connect("sqlite3", "testdata/test.db")
	if err != nil {
		t.Fatalf("Error connecting to db: #%v", err)
	}
	// 2. Check if migration metadata exists
	metadataQuery := "SELECT * FROM DATABASE_MIGRATION_LOGS ORDER BY installed_rank"
	expectedVersions := [2]string{"v001-schema.sql", "v001-data.sql"}

	rows, err := db.Query(metadataQuery)
	if err != nil {
		t.Fatalf("Error querying migration metadata: #%v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			t.Fatalf("Error closing migration metadata rows: #%v", err)
		}
	}()
	if !rows.Next() {
		t.Fatalf("No metadata found")
	}

	for _, v := range expectedVersions {
		var installedRank int
		var version string
		var query string
		var checksum uint32
		var installedOn string
		var executionTime int
		err = rows.Scan(&installedRank, &version, &query, &checksum, &installedOn, &executionTime)
		if err != nil {
			t.Fatalf("Error scanning metadata: #%v", err)
		}
		assert.Equal(t, v, version)
		rows.Next()
	}

	// 3. Check if data exists in user table
	userQuery := "SELECT COUNT(*) FROM USER"
	var count int
	err = db.QueryRow(userQuery).Scan(&count)
	if err != nil {
		t.Fatalf("Error querying user table: #%v", err)
	}
	assert.Equal(t, 1, count)
}
