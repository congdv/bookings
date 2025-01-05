package handlers

import (
	"net/http"

	"github.com/congdv/bookings/pkg/config"
	"github.com/congdv/bookings/pkg/models"
	"github.com/congdv/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// create new repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// set repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repository *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repository.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{StringMap: stringMap})
}

func (repository *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	remoteIP := repository.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
