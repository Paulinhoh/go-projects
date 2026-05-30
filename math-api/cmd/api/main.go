package main

import (
	"log"
	"math-api/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/sum", handlers.Sum)
	mux.HandleFunc("POST /api/sub", handlers.Sub)
	mux.HandleFunc("POST /api/mult", handlers.Mult)
	mux.HandleFunc("POST /api/div", handlers.Div)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
