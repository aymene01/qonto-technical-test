package main

import (
	"log"
	"net/http"
)


func main() {
	router := http.NewServeMux()
	loadRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

