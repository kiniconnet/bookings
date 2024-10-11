package handlers

import (
	"net/http"

	"github.com/kiniconnet/bookings/pkg/config"
	"github.com/kiniconnet/bookings/pkg/models"
	"github.com/kiniconnet/bookings/pkg/render"
)



var Repo *Repository

type Repository struct {
	App config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: *a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{

	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Perform some business logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
