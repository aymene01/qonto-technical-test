package main

import (
	"net/http"

	"github.com/aymene01/qonto-technical-test/internal/handlers"
)

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

	router.HandleFunc("POST /count-frequency", exercice.CountFrequencyHandler)
	router.HandleFunc("POST /is-truthy", exercice.IsTruthyHandler)
	router.HandleFunc("POST /is-student-of", exercice.IsStudentOfHandler)
	router.HandleFunc("POST /generate-even-int-list", exercice.GenerateEvenIntListHandler)
	router.HandleFunc("POST /parse-int", exercice.ParseIntHandler)
	router.HandleFunc("POST /was-student-during", exercice.WasStudentDuringHandler)
}