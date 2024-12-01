package config

type ServerConfig struct {
	Port int 	`yaml:"port"`
	Env string	`yaml:"env"`
	Version string	`yaml:"version"`
}