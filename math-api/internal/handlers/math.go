package handlers

import (
	"encoding/json"
	"math-api/internal/models"
	"math-api/internal/services"
	"net/http"
)

func Sum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var operacao models.Operacao
	if err := json.NewDecoder(r.Body).Decode(&operacao); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "body invalido",
		})
		return
	}

	sum := services.Sum(operacao.NumA, operacao.NumB)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]float64{
		"result": sum,
	})
}

func Sub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var operacao models.Operacao
	if err := json.NewDecoder(r.Body).Decode(&operacao); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "body invalido",
		})
		return
	}

	sub := services.Sub(operacao.NumA, operacao.NumB)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]float64{
		"result": sub,
	})
}

func Mult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var operacao models.Operacao
	if err := json.NewDecoder(r.Body).Decode(&operacao); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(r.Body)
		return
	}

	mult := services.Mult(operacao.NumA, operacao.NumB)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]float64{
		"result": mult,
	})
}

func Div(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var operacao models.Operacao
	if err := json.NewDecoder(r.Body).Decode(&operacao); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "body invalido",
		})
		return
	}

	div, err := services.Div(operacao.NumA, operacao.NumB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]float64{
		"result": div,
	})
}
