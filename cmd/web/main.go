package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jaysg1thub/bookings/internal/config"
	"github.com/jaysg1thub/bookings/internal/handlers"
	"github.com/jaysg1thub/bookings/internal/models"
	"github.com/jaysg1thub/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function:
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	fmt.Printf("Starting application on port %s", portNumber)

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// what am i going to put in the session?
	gob.Register(models.Reservation{})

	// chanbge this to true when in Production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // sessions last for 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // in Prod this will be "true"

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	// creates repository variable
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// gives render pkg access to the "config.AppConfig" info
	render.NewTemplates(&app)

	return nil
}
