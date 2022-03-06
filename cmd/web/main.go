package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mojafa/go-course/pkg/config"
	"github.com/mojafa/go-course/pkg/handlers"
	"github.com/mojafa/go-course/pkg/render"
)

const portNumber = ":8080"

var (
	app     config.AppConfig
	//varibale shadowing
	session *scs.SessionManager
)

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 3 * time.Hour
	// session.IdleTimeout = 20 * time.Minute
	// session.Cookie.Name = "session_id"
	// session.Cookie.Domain = "example.com"
	// session.Cookie.HttpOnly = true
	// session.Cookie.Path = "/example/"
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
