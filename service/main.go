package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TeddyCr/priceitt/service/infrastructure/database/postgres"
	"github.com/TeddyCr/priceitt/service/infrastructure/fernet"
	"github.com/TeddyCr/priceitt/service/infrastructure/jwt_secret"
	"github.com/TeddyCr/priceitt/service/models/config"
	"github.com/TeddyCr/priceitt/service/resource"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	config := getConfig()

	logger := getLoggerConfig(config)

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(60 * time.Second))

	fernet.Initialize(config.Fernet)
	jwt_secret.Initialize(config.JwTokenConfiguration)
	mountRoutes(r, config)

	if config.Server.Type == "http" {
		err := http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), r)
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServeTLS(":"+strconv.Itoa(config.Server.Port), config.Server.Certificate, config.Server.Key, r)
		if err != nil {
			panic(err)
		}
	}
}

func mountRoutes(r chi.Router, config config.Config) {
	pq, err := postgres.PersistencePostgres{}.Initialize(config.Database)
	if err != nil {
		panic(err)
	}
	r.Mount("/api/v1/auth", resource.NewAuthResource(pq).Routes())
	r.Mount("/api/v1/user", resource.NewUserResource(pq).Routes())
}

func getLoggerConfig(config config.Config) *httplog.Logger {
	return httplog.NewLogger("service", httplog.Options{
		LogLevel: slog.LevelDebug,
		JSON:     true,
		Concise:  true,
		Tags: map[string]string{
			"env":     config.Logging.Level,
			"version": config.Server.Version,
		},
		TimeFieldFormat: time.RFC3339,
		RequestHeaders:  true,
		ResponseHeaders: true,
	})
}

func getConfig() config.Config {
	path := getConfigFilePath()
	raw_confg, err := os.ReadFile(path)
	mapper := getEnvVarMapper()
	if err != nil {
		panic(err)
	}

	expanded_config := []byte(os.Expand(string(raw_confg), mapper))

	var config config.Config
	err = yaml.Unmarshal(expanded_config, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func getConfigFilePath() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(root)

	var val string
	var found bool
	val, found = os.LookupEnv("SERVER_CONFIG_FILE_PATH")
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
