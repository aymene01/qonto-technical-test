package main

import (
	"log"
	"net/http"

	handlers "github.com/aymene01/qonto-technical-test/internal/handlers"
)


func main() {
	router := http.NewServeMux()
	loadRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func loadRoutes(router *http.ServeMux){
	loadHealthRoutes(router)
	loadExerciceRoutes(router)
}

func loadHealthRoutes(router *http.ServeMux){
	health := handlers.Health{}
	router.HandleFunc("/health", health.GetHealthStatus)
}

func loadExerciceRoutes(router *http.ServeMux){
	exercices := handlers.Exercice{}

	router.HandleFunc("POST /countCharFrequency", exercices.CountFrequencyHandler)
	router.HandleFunc("POST /isTruthy", exercices.IsTruthyHandler)
	router.HandleFunc("POST /isStudentOf", exercices.IsStudentOfHandler)
	router.HandleFunc("POST /generateEvenIntList", exercices.GenerateEvenIntListHandler)
	router.HandleFunc("POST /parseInt", exercices.ParseIntHandler)
}
