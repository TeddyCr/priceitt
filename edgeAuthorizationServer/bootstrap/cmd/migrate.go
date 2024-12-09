package main

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"priceitt.xyz/edgeAuthorizationServer/models"
	"github.com/TeddyCr/priceitt/utils/migrations"
)

func run() {
	config := getConfigModel()
	migrations.ExecMigration(config.Migration, config.Database)
}

func getConfigModel() models.Config {
	path := getConfigFile()
	raw_confg, err := os.ReadFile(path)
	mapper := getEnvVarMapper()
	if err != nil {
		panic(err)
	}

	expanded_config := []byte(os.Expand(string(raw_confg), mapper))

	var config models.Config
	err = yaml.Unmarshal(expanded_config, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func getConfigFile() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(root)

	var val string
	var found bool
	val, found = os.LookupEnv("EDGE_AUTHORIZATION_SERVER_CONFIG_FILE_PATH")
	if !found {
		val = "config/config.yaml"
	}

	return filepath.Join(parent, val)
}

func getEnvVarMapper() func(string) string {
	mapper := func(envName string) string {
		split := strings.Split(envName, ":-")
		defaultValue := ""
		if len(split) == 2 {
			envName = split[0]
			defaultValue = split[1]
		}

		val, ok := os.LookupEnv(envName)
		if !ok {
			// Postgres uses $ as a placeholder for parameters i.e `$1, $2, etc.` 
			if len(defaultValue) < 1 {
				return "$" + envName
			}
			return defaultValue
		}
		return val
	}

	return mapper
}

func main() {
	run()
}