package models

type MigrationMetadata struct {
	Version       string
	Query         string
	Checksum      uint32
	ExecutionTime int
}

type MigrationMetadataQueries struct {
	InsertMetadataQuery           string `yaml:"insertMetadataQuery"`
	SelectMigrationVersionsQuery  string `yaml:"selectMigrationVersionsQuery"`
	SelectMigrationChecksumsQuery string `yaml:"selectMigrationChecksumsQuery"`
}

type MigrationConfig struct {
	SchemaPath      string                   `yaml:"schemaPath"`
	DataPath        string                   `yaml:"dataPath"`
	MetadataPath    string                   `yaml:"metadataPath"`
	MetadataQueries MigrationMetadataQueries `yaml:"metadataQueries"`
	Force           bool                     `yaml:"force"`
	CheckIntegrity  bool                     `yaml:"checkIntegrity"`
}
