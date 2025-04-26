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
	id := chi.URLParam(r, "category_id)")

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
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}
	fmt.Println("Product data:", product)
	// fmt.Println("Product data category id :", product.CategoryID)
	// fmt.Println("Product data device type :", product.DeviceType)
	// fmt.Println("Product data model :", product.Model)
	// fmt.Println("Product data color :", product.Color)
	// fmt.Println("Product data storage :", product.Storage)
	// fmt.Println("Product data battery health :", product.BatteryHealth)
	// fmt.Println("Product data processor :", product.Processor)
	// fmt.Println("Product data ram :", product.Ram)
	// fmt.Println("Product data description :", product.Description)
	// fmt.Println("Product data created at :", product.CreatedAt)

	if err := ProductStore.Add(product); err != nil {
		fmt.Println("Error creating product:", err)
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
