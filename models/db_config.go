package models

type DatabaseConfig struct {
	DriverClass      string `yaml:"driverClass"`
	ConnectionString string `yaml:"connectionString"`
}
