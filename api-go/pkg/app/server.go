package app

import "github.com/gorilla/mux"

// App is the server structure.
type App struct {
	Router *mux.Router
}

// New func that will be in charge of returning the actual application based on our struct.
func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")

	a.Router.HandleFunc("/detect-labels", a.DetectLabelsHandler()).Methods("POST")
	a.Router.HandleFunc("/compare-faces", a.CompareFacesHandler()).Methods("POST")
}
