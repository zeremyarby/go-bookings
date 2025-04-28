package main

import (
	"net/http"

	"github.com/zeremyarby/go-bookings/pkg/config"
	"github.com/zeremyarby/go-bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Using chi middleware
	mux.Use(middleware.Recoverer)

	mux.Use(NoSurf)      // CSRF protection
	mux.Use(SessionLoad) // Session management

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
