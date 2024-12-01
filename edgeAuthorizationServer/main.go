package main

import (
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"priceitt.xyz/edgeAuthorizationServer/resource"
)

// WHat do I need to do here?
// 1. I need to register a repository. We should have a repository package.
// 2. I need to register my resources.
// 3. I need model my entities, create my serializers, and create my handlers (business logic).


func main() {
	logger := getLoggerConfig()

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
  	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/user", resource.NewUserResource(nil).Routes())
}

func getLoggerConfig() *httplog.Logger {
	return httplog.NewLogger("edgeAuthorizationServer", httplog.Options{
		LogLevel: slog.LevelDebug,
		JSON: true,
		Concise: true,
		Tags: map[string]string{
			"env": "dev",
			"version": "1.0.0",
		},
		TimeFieldFormat:  time.RFC3339,
		RequestHeaders: true,
		ResponseHeaders: true,
	})
}