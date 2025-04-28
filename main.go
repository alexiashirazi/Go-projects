package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	// Backend
	"curs1_boilerplate/cmd/backend/handler"
	"curs1_boilerplate/cmd/backend/store"

	// Frontend (Templ)
	"curs1_boilerplate/cmd/frontend/views/base"
	"curs1_boilerplate/cmd/frontend/views/pages"
)

/*
wgo -file '\.templ$' templ generate & wgo -file .go go run .
*/
func main() {
	// DB Connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(" Conexiune DB eșuată:", err)
	}
	defer conn.Close(ctx)

	handler.CategoryStore = store.NewDbCategoryStore(conn)
	handler.UserStore = store.NewDbUserStore(conn)
	handler.ProductStore = store.NewDbProductStore(conn)

	// Router
	r := chi.NewRouter()

	// --- FRONTEND ROUTES ---
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			log.Println("No user_id cookie:", err)
		}

		isLogged := cookie != nil && cookie.Value != ""

		base.PageSkeleton(pages.MainPage(isLogged), isLogged).Render(r.Context(), w)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		base.PageSkeleton(pages.Login(), false).Render(r.Context(), w)
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		base.PageSkeleton(pages.Register(), false).Render(r.Context(), w)
	})
	r.Get("/sell", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			log.Println("error cookie sell", err)
		}
		isLogged := cookie != nil && cookie.Value != ""
		if !isLogged {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		base.PageSkeleton(pages.Sell(), isLogged).Render(r.Context(), w)

	})

	fs := http.StripPrefix("/images/", http.FileServer(http.Dir("./cmd/frontend/views/images/public/images")))
	r.Handle("/images/*", fs)

	// --- BACKEND ROUTES ---
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)

	r.Group(func(r chi.Router) {
		r.Use(handler.RequireAuth)
		r.Get("/profile", handler.Profile)
		r.Get("/logout", handler.Logout)
		r.Post("/api/products", handler.CreateProduct)
		r.Get("/sell/form", func(w http.ResponseWriter, r *http.Request) {
			categoryID := r.URL.Query().Get("category_id")
			if categoryID == "" {
				http.Error(w, "Missing category_id", http.StatusBadRequest)
				return
			}

			// APELI CATEGORIA
			category, err := handler.CategoryStore.GetCategoryByID(categoryID)
			if err != nil {
				http.Error(w, "Category not found", http.StatusNotFound)
				return
			}

			// Dai mai departe la pagină categoria
			base.PageSkeleton(pages.SellForm(category.ID, category.Name), true).Render(r.Context(), w)
		})

	})

	r.Get("/api/products", func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("your_cookie_name") // presupun că iei login info
		isLogged := cookie != nil && cookie.Value != ""

		products, err := handler.GetAllProducts2(r)
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}

		// Acum ai "products" și îl poți trimite
		base.PageSkeleton(pages.Products(products), isLogged).Render(r.Context(), w)
	})
	r.Get("/api/products/{id}", handler.GetProductByID)
	r.Get("/api/products/category/{categoryID}", handler.GetProductsByCategory)
	r.Put("/api/products/{id}", handler.UpdateProduct)
	r.Delete("/api/products/{id}", handler.DeleteProduct)
	r.Get("/api/categories", handler.GetAllCategories)
	r.Get("/api/categories/{id}", handler.GetCategoryByID)
	r.Post("/api/categories", handler.CreateCategory)
	r.Put("/api/categories/{id}", handler.UpdateCategory)
	r.Delete("/api/categories/{id}", handler.DeleteCategory)
	r.Get("/api/products/user/{userID}", handler.GetProductsByUserID)

	// Start server
	log.Println("Server pornit pe http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
