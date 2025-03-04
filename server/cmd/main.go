package main

import (
	"log"
	"net/http"

	"github.com/magomzr/archium/api"
)

func main() {
	r := api.SetupRoutes()

	port := ":8080"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
