package main

import (
	"green-api/routes"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
