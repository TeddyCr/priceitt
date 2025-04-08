package config

import (
	"github.com/TeddyCr/priceitt/service/models"
)

type Config struct {
	Server               ServerConfig           `yaml:"server"`
	Fernet               FernetConfig           `yaml:"fernet"`
	Logging              LoggingConfig          `yaml:"logging"`
	Database             models.DatabaseConfig  `yaml:"database"`
	JwTokenConfiguration SecretConfig           `yaml:"jwTokenConfiguration"`
	Migration            models.MigrationConfig `yaml:"migration"`
}
