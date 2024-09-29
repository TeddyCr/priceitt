package migrations

import (
	"database/sql"
	"flag"
	"hash/crc32"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/utils/database"
	"github.com/TeddyCr/priceitt/utils/files"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func ExecMigration(migrationConfig models.MigrationConfig, dbConfig models.DatabaseConfig) {
	force := migrationConfig.Force
	checkIntegrity := migrationConfig.CheckIntegrity
	flag.Parse()

	migrate := Migrate{
		force:           force,
		checkIntegrity:  checkIntegrity,
		db:              database.Connect(dbConfig),
		migrationConfig: migrationConfig,
		fileChecksum:    make(map[string]uint32),
	}
	migrate.createMigrationMetadataTable()
	migrate.migrate()
}

type Migrate struct {
	force           bool
	checkIntegrity  bool
	db              *sqlx.DB
	migrationConfig models.MigrationConfig
	fileChecksum    map[string]uint32
}

func (m *Migrate) migrate() {
	m.migrateSchema()
	m.migrateData()
}

func (m *Migrate) migrateSchema() {
	migrationFiles := m.getMigrationFiles(m.migrationConfig.SchemaPath)
	m.migrateFiles(migrationFiles)
}

func (m *Migrate) migrateData() {
	migrationFiles := m.getMigrationFiles(m.migrationConfig.DataPath)
	m.migrateFiles(migrationFiles)
}

func (m *Migrate) migrateFiles(migrationFiles []string) {
	migrationMetadata := m.getMigrationVersions()
	for _, migrationFile := range migrationFiles {
		version := filepath.Base(migrationFile)
		queries := m.getQueries(migrationFile, version)
		if !slices.Contains(migrationMetadata, version) || m.force {
			log.Printf("Running migration: %s", version)
			m.runMigration(queries, version)
		} else {
			log.Printf("Migration %s already executed", version)
		}
	}
}

func (m *Migrate) getMigrationFiles(path string) []string {
	root := files.GetRoot()
	migrationFiles, err := filepath.Glob(root + path + "/*.sql")
	if err != nil {
		log.Fatalf("Error retrieving migration files: #%v", err)
	}
	slices.Sort(migrationFiles)
	return migrationFiles
}

func (m *Migrate) getQueries(migrationFile string, version string) []string {
	queriesBytes, err := os.ReadFile(migrationFile)
	m.fileChecksum[version] = crc32.ChecksumIEEE(queriesBytes)
	if err != nil {
		log.Fatalf("Error reading migration file: #%v", err)
	}
	queriesStr := string(queriesBytes)
	queries := strings.Split(queriesStr, "\n\n")
	return queries
}

func (m *Migrate) runMigration(queries []string, version string) {
	if m.checkIntegrity {
		log.Printf("Checking query migration integrity for version: %s", version)
		migrationChecksum := m.getMigrationChecksum(version)
		if m.fileChecksum[version] != migrationChecksum && migrationChecksum != 0 {
			log.Fatalf(
				"Checksum mismatch for version %s.\nDB checksum: %d\nFile checksum: %d",
				version, migrationChecksum, m.fileChecksum[version],
			)
		}
	}
	for _, query := range queries {
		start := time.Now()
		_, err := m.db.Exec(query)
		elapsed := time.Since(start)
		if err != nil {
			log.Fatalf("Error running migration: #%v", err)
		}
		log.Printf("Migration %s executed in %s", version, elapsed)
		migrationMetadata := models.MigrationMetadata{
			Version:       version,
			Query:         query,
			Checksum:      m.fileChecksum[version],
			ExecutionTime: int(elapsed.Microseconds()),
		}
		m.writeMigrationMetadata(migrationMetadata)
	}
}

func (m *Migrate) getMigrationVersions() []string {
	var version []string
	query := m.migrationConfig.MetadataQueries.SelectMigrationVersionsQuery
	var rows *sql.Rows
	var err error

	rows, err = m.db.Query(query)
	if err != nil {
		log.Fatalf("Error querying migration metadata: #%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var migration string
		err = rows.Scan(&migration)
		if err != nil {
			log.Fatalf("Databse error: #%v", err)
		}
		version = append(version, migration)
	}
	return version
}

func (m *Migrate) getMigrationChecksum(version string) uint32 {
	var checksums []uint32
	query := m.migrationConfig.MetadataQueries.SelectMigrationChecksumsQuery + ` WHERE version = $1 LIMIT 1`
	var rows *sql.Rows
	var err error

	rows, err = m.db.Query(query, version)
	if err != nil {
		log.Fatalf("Error querying migration metadata: #%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var checksum uint32
		err = rows.Scan(&checksum)
		if err != nil {
			log.Fatalf("Databse error: #%v", err)
		}
		checksums = append(checksums, checksum)
	}

	if len(checksums) == 0 {
		return 0
	}
	return checksums[0]
}

func (m *Migrate) writeMigrationMetadata(migrationMetadata models.MigrationMetadata) {
	// Check if migration already exists
	var exists bool
	exists = true
	var query string
	query = m.migrationConfig.MetadataQueries.SelectMigrationChecksumsQuery + ` WHERE query = $1`
	var rows *sql.Rows
	var err error

	rows, err = m.db.Query(query, migrationMetadata.Query)
	if err != nil {
		log.Fatalf("Error querying migration metadata: #%v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		exists = false
	}

	if !exists {
		query = m.migrationConfig.MetadataQueries.InsertMetadataQuery // `INSERT INTO "DATABASE_MIGRATION_LOGS" (version, query, checksum, execution_time) VALUES ($1, $2, $3, $4)`
		_, err := m.db.Exec(query, migrationMetadata.Version, migrationMetadata.Query, migrationMetadata.Checksum, migrationMetadata.ExecutionTime)
		if err != nil {
			log.Fatalf("Error writing migration metadata: #%v", err)
		}
	}
}

func (m *Migrate) createMigrationMetadataTable() {
	path := m.getMigrationFiles(m.migrationConfig.MetadataPath)
	queries := m.getQueries(path[0], "metadata")
	log.Println("Creating migration metadata table if not exists.")
	for _, query := range queries {
		_, err := m.db.Exec(query)
		if err != nil {
			log.Fatalf("Error creating metadata table: #%v", err)
		}
	}
}
