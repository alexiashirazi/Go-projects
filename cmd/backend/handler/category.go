package handler

import (
	"curs1_boilerplate/cmd/backend/model"
	"curs1_boilerplate/cmd/backend/store"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var CategoryStore store.CategoryStore

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := CategoryStore.GetAllCategories()
	if err != nil {
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing category ID", http.StatusBadRequest)
		return
	}

	category, err := CategoryStore.GetCategoryByID(id)
	if err != nil {
		http.Error(w, "Error fetching category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid category data", http.StatusBadRequest)
		return
	}

	if err := CategoryStore.AddCategory(category); err != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing category ID", http.StatusBadRequest)
		return
	}

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid category data", http.StatusBadRequest)
		return
	}

	category.ID = id

	if err := CategoryStore.UpdateCategory(category); err != nil {
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing category ID", http.StatusBadRequest)
		return
	}

	if err := CategoryStore.DeleteCategory(id); err != nil {
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
