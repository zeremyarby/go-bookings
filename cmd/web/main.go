package main

import (
	"fmt"
	"log"
	"myapp3/pkg/config"
	"myapp3/pkg/handlers"
	"myapp3/pkg/renders"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main application function
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 30 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // In production, set to true

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	renders.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)

	fmt.Println("Starting application on port 8080")
	// fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(":8080", nil)
}
