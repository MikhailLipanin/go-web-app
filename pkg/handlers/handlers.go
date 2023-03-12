package handlers

import (
	"github.com/MikhailLipanin/how2amuse/pkg/config"
	"github.com/MikhailLipanin/how2amuse/pkg/driver"
	"github.com/MikhailLipanin/how2amuse/pkg/models"
	"github.com/MikhailLipanin/how2amuse/pkg/render"
	"log"
	"net/http"
	"strconv"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  *driver.DB
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  db,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Catalog page handler
func (m *Repository) Catalog(w http.ResponseWriter, r *http.Request) {
	regions, err := m.DB.GetRegions()
	if err != nil {
		log.Fatal(err)
	}
	cities, err := m.DB.GetCities()
	if err != nil {
		log.Fatal(err)
	}
	// send the data to the template
	render.RenderTemplate(w, "catalog.page.html", &models.TemplateData{
		Regions: regions,
		Cities:  cities,
	})

}

// City page handler
func (m *Repository) City(w http.ResponseWriter, r *http.Request) {
	city, err := strconv.Atoi(r.URL.Query().Get("city"))
	if err != nil {
		log.Fatal(err)
	}
	places, err := m.DB.GetPlaces()
	if err != nil {
		log.Fatal(err)
	}
	// send the data to the template
	render.RenderTemplate(w, "city.page.html", &models.TemplateData{
		CityID: city,
		Places: places,
	})
}
