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

	migrations.ExecMigration(config.Migration, config.Database)
}

func getConfigFile() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(root)
	return filepath.Join(parent, "config/config.yaml")
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