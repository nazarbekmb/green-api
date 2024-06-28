package routes

import (
	"green-api/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.IndexHandler)
}
