package main

import (
	"fmt"
	"github.com/MikhailLipanin/how2amuse/pkg/config"
	"github.com/MikhailLipanin/how2amuse/pkg/driver"
	"github.com/MikhailLipanin/how2amuse/pkg/handlers"
	"github.com/MikhailLipanin/how2amuse/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=places_db user=postgres password=1234")
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
	log.Println("Connected to database!")
	defer db.SQL.Close()

	// create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache!")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	log.Println(fmt.Sprintf("Starting Application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
