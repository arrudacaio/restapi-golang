// This package will represent the structure of our app
// There will be two fields, our app is going to have a DB and a Router.
package app

import "github.com/gorilla/mux"

// Question: What *mux.Router means?
type App struct {
	Router *mux.Router
}
