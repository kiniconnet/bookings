package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kiniconnet/bookings/pkg/config"
	"github.com/kiniconnet/bookings/pkg/handlers"
	"github.com/kiniconnet/bookings/pkg/render"
)

const PortNumber = ":4000"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true in production

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction


	app.Session = session

	tc, err := render.CreateTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	err = render.NewTemplate(&app)
	if err != nil {
		log.Fatal("could not render the template")
		return
	}

	
	fmt.Println("The server is running in port ", PortNumber)

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err !=  nil {
		log.Fatal(err)
	}
}
