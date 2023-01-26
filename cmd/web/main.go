package main

import (
	"fmt"
	"github.com/MikhailLipanin/go-web-app/pkg/config"
	"github.com/MikhailLipanin/go-web-app/pkg/handlers"
	"github.com/MikhailLipanin/go-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
