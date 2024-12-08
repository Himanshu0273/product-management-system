package router

import (
    "github.com/gorilla/mux"
    "product-management-system/api/handlers" // Correct import for handlers
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
    r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
    return r
}
