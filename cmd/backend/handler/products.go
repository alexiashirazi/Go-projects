package handler

import (
	"curs1_boilerplate/cmd/backend/model"
	"curs1_boilerplate/cmd/backend/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var ProductStore store.ProductStore

// GetAllProducts handles the request to get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ProductStore.GetAll()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetAllProducts2(r *http.Request) ([]model.Product, error) {
	products, err := ProductStore.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID handles the request to get a product by its ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	product, err := ProductStore.GetByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProductsByCategory handles the request to get products by category
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "category_id")

	products, err := ProductStore.GetByCategory(id)
	if err != nil {
		http.Error(w, "Error fetching products by category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateProduct handles the request to create a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println(product)
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	// SCOȚI userID din context (cookie)
	userID, ok := r.Context().Value("userID").(string)

	if !ok || userID == "" {
		fmt.Println(userID)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Legi produsul de utilizator
	product.UserID = userID

	if err := ProductStore.Add(product); err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
}

// UpdateProduct handles the request to update an existing product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	// Ensure ID in URL matches product ID
	product.ID = id

	if err := ProductStore.Update(product); err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

// DeleteProduct handles the request to delete a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := ProductStore.Delete(id); err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

// GetProductsByCategory handles the request to get products by category
func GetProductsByUserID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")

	products, err := ProductStore.GetByUser(id)
	if err != nil {
		http.Error(w, "Error fetching products by category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
