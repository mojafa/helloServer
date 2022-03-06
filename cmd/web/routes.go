package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mojafa/go-course/pkg/config"
	"github.com/mojafa/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux

	r := chi.NewRouter()
	// r.Use(WriteToConsole)
	r.Use(NoSurf)
	r.Use(SessionLoad) 
	r.Use(middleware.Recoverer)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	return r
}
