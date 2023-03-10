package handlers

import (
	"github.com/MikhailLipanin/how2amuse/pkg/config"
	"github.com/MikhailLipanin/how2amuse/pkg/driver"
	"github.com/MikhailLipanin/how2amuse/pkg/models"
	"github.com/MikhailLipanin/how2amuse/pkg/render"
	"log"
	"net/http"
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
	data := make(map[string]any)
	reg, err := m.DB.GetRegions()
	if err != nil {
		log.Fatal(err)
	}
	data["test"] = reg

	// send the data to the template
	render.RenderTemplate(w, "catalog.page.html", &models.TemplateData{
		Data: data,
	})
}
