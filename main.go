package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Author string `json:"author"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("COntent-type", "application/json")
		w.WriteHeader(http.StatusOK)

		data := Response{
			Status: "success",
			Message: "Olá! Bem-vindo a minha primeira API em GO",
			Author: "Vitor H.",
		}

		json.NewEncoder(w).Encode(data)
	})

	println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}