package config

type SecretConfig struct {
	Secret  	string `yaml:"secret"`
	Issuer 		string `yaml:"issuer"`
	Audience 	string `yaml:"audience"`
}
