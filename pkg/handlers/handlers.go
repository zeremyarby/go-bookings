package handlers

import (
	"net/http"

	"github.com/zeremyarby/go-bookings/pkg/config"
	"github.com/zeremyarby/go-bookings/pkg/models"
	"github.com/zeremyarby/go-bookings/pkg/renders"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home this is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About this is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send the data to the template
	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
