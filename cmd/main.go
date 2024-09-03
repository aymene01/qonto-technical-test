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
	exercice := handlers.Exercice{}

	router.HandleFunc("POST /countCharFrequency", exercice.CountFrequencyHandler)
	router.HandleFunc("POST /isTruthy", exercice.IsTruthyHandler)
	router.HandleFunc("POST /isStudentOf", exercice.IsStudentOfHandler)
	router.HandleFunc("POST /generateEvenIntList", exercice.GenerateEvenIntListHandler)
	router.HandleFunc("POST /parseInt", exercice.ParseIntHandler)
}
