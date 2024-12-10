package config

type ServerConfig struct {
	Type string	`yaml:"type"`
	Port int 	`yaml:"port"`
	Env string	`yaml:"env"`
	Version string	`yaml:"version"`
	Certificate string	`yaml:"certificate"`
	Key string	`yaml:"key"`
}