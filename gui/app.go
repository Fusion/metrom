package main

import (
	"context"
	"gui/components"
	"gui/net"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func NewChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//r.Get("/closemodal", closeModal)
	r.Mount("/hops", HopsCollector())

	r.Get("/initial", templ.Handler(components.Pages([]struct {
		Path  string
		Label string
	}{
		{"/greet", "Greet form"},
	}, struct {
		Version string
		Text    string
	}{
		version, "No update available",
	})).ServeHTTP)
	r.Get("/closemodal", closeModal)

	r.Get("/greet", templ.Handler(components.GreetForm("/greet")).ServeHTTP)
	r.Post("/greet", components.Greet)
	r.Get("/modal", templ.Handler(components.TestPage("#modal", "outerHTML")).ServeHTTP)
	r.Post("/modal", templ.Handler(components.ModalPreview("Title for the modal", "Sample Data")).ServeHTTP)
	return r
}

func openModal(w http.ResponseWriter, r *http.Request, title string, message string) {
	w.Header().Add("Cache-Control", "no-store")
	w.Header().Set("HX-Retarget", "#modalwrapper")
	component := components.Modal(title, message)
	component.Render(r.Context(), w)
}

func closeModal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	// Never write an empty string; wail's middleware makes some assumptions
	w.Write(nil)
}

func HopsCollector() chi.Router {
	r := chi.NewRouter()
	netController := net.NewNetControllerV2()
	netController.SetSetting("resolve", "off")
	resolver := net.NewResolver()
	hopsCollection := NewHopsCollection(&netController, resolver)
	r.Get("/", hopsCollection.Get)
	r.Post("/", hopsCollection.Post)
	r.Post("/toggleresolve", hopsCollection.ToggleResolve)
	r.Get("/resetsearch", hopsCollection.ResetSearch)
	return r
}
