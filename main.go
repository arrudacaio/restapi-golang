package main

import (
	"log"
	"net/http"
	"os"
	"restapi/app"
	"restapi/app/database"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running...")
	err = http.ListenAndServe(":9000", nil)
	check(err)

}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
