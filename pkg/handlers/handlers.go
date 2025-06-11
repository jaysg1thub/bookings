package handlers

import (
	"net/http"

	"github.com/jaysg1thub/bookings/pkg/config"
	"github.com/jaysg1thub/bookings/pkg/models"
	"github.com/jaysg1thub/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repsotiry is the repository type
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

// Home:	a handler function for the HomePage:
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// grab the remote IP address & store in session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About:	is the about page handler:
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// pull IP value out of session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
