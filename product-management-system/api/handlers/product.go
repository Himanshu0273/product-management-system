package handlers

import (
    "encoding/json"
    "net/http"
    "product-management-system/api/models"   // Correct import for models
    "product-management-system/api/database" // Correct import for database
    "github.com/gorilla/mux"
)

// CreateProduct creates a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Save the product in PostgreSQL
    err = database.SaveProduct(&product)
    if err != nil {
        http.Error(w, "Failed to create product", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// GetProductByID retrieves a product by its ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    product, err := database.GetProductByID(id)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(product)
}

// GetAllProducts retrieves all products for a given user with filtering
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    products, err := database.GetAllProductsByUser(userID)
    if err != nil {
        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(products)
}
