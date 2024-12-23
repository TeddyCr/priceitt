package models

import (
	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/models/config"
)



type Config struct {
	Server		config.ServerConfig `yaml:"server"`
	Fernet 		config.FernetConfig `yaml:"fernet"`
	Logging		config.LoggingConfig `yaml:"logging"`
	Database	models.DatabaseConfig `yaml:"database"`
	Migration	models.MigrationConfig `yaml:"migration"`
}