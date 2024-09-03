package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from the api #CI/CD"))
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
