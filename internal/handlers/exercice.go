package handlers

import (
	"encoding/json"
	"net/http"

	f "github.com/aymene01/qonto-technical-test/internal/functions"
	"github.com/aymene01/qonto-technical-test/internal/utils"
)

type Exercice struct {}

// -> 
func (e *Exercice) CountFrequencyHandler(w http.ResponseWriter, r *http.Request) {
	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f.CountFrequency(input)
	utils.RespondJSON(w, http.StatusOK, result)
}

// -> 
func (e *Exercice) IsStudentOfHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Obj     f.ObjArg `json:"obj"`
		College string `json:"college"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f.IsStudentOf(input.Obj, input.College)
	response := map[string]bool{"result": result}
	utils.RespondJSON(w, http.StatusOK, response)
}

// ->
func (e *Exercice) IsTruthyHandler(w http.ResponseWriter, r *http.Request) {
	var input []any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f.IsTruthy(input)
	response := map[string]bool{"result": result}
    utils.RespondJSON(w, http.StatusOK, response)
}

// ->
func (e *Exercice) GenerateEvenIntListHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        N int `json:"n"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := f.GenerateEvenIntList(input.N)
    response := map[string][]int{"result": result}
	utils.RespondJSON(w, http.StatusOK, response)
}

//
func (e *Exercice) ParseIntHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Value string `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := f.ParseInt(input.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]int{"result": result}
	utils.RespondJSON(w, http.StatusOK, response)
}
