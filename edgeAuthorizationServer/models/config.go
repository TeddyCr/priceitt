package models

import (
	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/models/config"
)



type Config struct {
	Server		config.ServerConfig `yaml:"server"`
	Logging		config.LoggingConfig `yaml:"logging"`
	Database	models.DatabaseConfig `yaml:"database"`
	Migration	models.MigrationConfig `yaml:"migration"`
}