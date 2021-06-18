// This package will represent the structure of our app
// There will be two fields, our app is going to have a DB and a Router.
package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", nil).Methods("GET")
}
