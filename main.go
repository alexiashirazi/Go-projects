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
		if err == nil {
			log.Println("error cookie sell", err)
		}
		isLogged := cookie != nil && cookie.Value != ""
		if !isLogged {
			base.PageSkeleton(pages.Login(), false).Render(r.Context(), w)
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
	})

	r.Get("/api/products", handler.GetAllProducts)
	r.Get("/api/products/{id}", handler.GetProductByID)
	r.Get("/api/products/category/{categoryID}", handler.GetProductsByCategory)
	r.Post("/api/products", handler.CreateProduct)
	r.Put("/api/products/{id}", handler.UpdateProduct)
	r.Delete("/api/products/{id}", handler.DeleteProduct)
	r.Get("/api/categories", handler.GetAllCategories)
	r.Get("/api/categories/{id}", handler.GetCategoryByID)
	r.Post("/api/categories", handler.CreateCategory)
	r.Put("/api/categories/{id}", handler.UpdateCategory)
	r.Delete("/api/categories/{id}", handler.DeleteCategory)

	// Start server
	log.Println("Server pornit pe http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
