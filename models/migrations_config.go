package models

type MigrationMetadata struct {
	Version       string
	Query         string
	Checksum      uint32
	ExecutionTime int
}

type MigrationMetadataQueries struct {
	InsertMetadataQuery string
	SelectMigrationVersionsQuery string
	SelectMigrationChecksumsQuery string
}

type MigrationConfig struct {
	SchemaPath string
	DataPath   string
	MetadataPath string
	MetadataQueries MigrationMetadataQueries
	Force 	   bool
	CheckIntegrity bool
}