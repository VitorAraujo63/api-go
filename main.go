package main

import (
	"log"
	"minha-api-go/handlers"
	"minha-api-go/repository"
	"net/http"
)

func main() {
	repo := repository.NewProductRepository()

	productHandler := handlers.NewProductHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products", productHandler.GetProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	log.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}