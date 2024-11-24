package main

import (
	"backend-test/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
