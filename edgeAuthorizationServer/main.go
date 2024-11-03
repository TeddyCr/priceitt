package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// WHat do I need to do here?
// 1. I need to register a repository. We should have a repository package.
// 2. I need to register my resources.
// 3. I need model my entities, create my serializers, and create my handlers (business logic).


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}