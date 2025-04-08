package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// 3. Auction REST API
// Problem:
// Create a basic HTTP REST API that simulates an auction website. You don’t need any frontend, just API routes that can be tested with curl, Postman, or a browser. Use only in-memory storage (e.g., a global slice or map).

// Each product should have:

// Name (string)
// Description (string)
// Price (float64)
// IsSold (bool)
// Required API endpoints:

// Add a product – POST /products
// Request body: JSON { "name": "Laptop", "description": "Gaming laptop", "price": 999.99 }
// Remove a product – DELETE /products/{name}
// Sell a product – POST /products/{name}/sell
// List all products – GET /products
// List only sold products – GET /products?sold=true
// General product view – GET /products
// Response: list of all products showing only name and price
// Detailed product view – GET /products/{name}
// Response: includes name, price, description, isSold
// Skills practiced:

// Creating an HTTP server with net/http
// REST API design
// Structs, slices, basic state management
// JSON marshaling/unmarshaling
// Route handling and URL parameters

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsSold      bool    `json:"isSold"`
}

var products []Product

func main() {
	router := chi.NewRouter()
	router.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "cant read from body", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, &product)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		products = append(products, product)
		marshaled, _ := json.Marshal(product)

		response := fmt.Sprintf("Got body %s", marshaled)
		w.Write([]byte(response))

	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		soldOnly := r.URL.Query().Get("sold") == "true"

		var prod []Product
		for _, p := range products {
			if soldOnly {
				if p.IsSold {
					prod = append(prod, p)
				}
			} else {
				prod = append(prod, p)
			}
		}

		resp, _ := json.Marshal(prod)
		response := fmt.Sprintf("Get body %s\n", resp)
		w.Write([]byte(response))

	})

	router.Delete("/products/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		found := false
		newProducts := []Product{}
		for _, p := range products {
			if p.Name == name {
				found = true
				continue
			}
			newProducts = append(newProducts, p)
		}

		if !found {
			http.Error(w, "product not found", http.StatusInternalServerError)
			return
		}
		products = newProducts
		w.Write([]byte("stergere efectuata"))
	})

	router.Post("/product/{name}/sell", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		found := false
		newProducts := []Product{}
		for _, p := range products {
			if p.Name == name {
				found = true
				if p.IsSold == true {
					http.Error(w, "product already sold ", http.StatusInternalServerError)
					return
				}
				p.IsSold = true
			}
			newProducts = append(newProducts, p)
		}
		if !found {
			http.Error(w, "product not found", http.StatusInternalServerError)
			return
		}
		products = newProducts
		w.Write([]byte("vanzare efectuata"))

	})

	router.Get("/products/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		for _, p := range products {
			if p.Name == name {
				jsonData, err := json.Marshal(p)
				if err != nil {
					http.Error(w, "couldnt get product", http.StatusInternalServerError)
					return
				}

				w.Write(jsonData)
				return
			}
		}

		http.Error(w, "product not found", http.StatusNotFound)
	})

	fmt.Println("Server pornit pe http://localhost:3000")
	http.ListenAndServe(":3000", router)

	return
}
