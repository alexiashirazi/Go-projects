package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5" //import la chi
)

type Todo struct {
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Deadline string `json:"deadline"`
}

// Middle ware care logheaza fiecare request
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request", r.Method, r.URL.Path) //afiseaza in terminal requestul si de unde vine
		next.ServeHTTP(w, r)                                  //trimite request ul la urmatorul handler
	})
}

func main() {

	router := chi.NewRouter()

	//Routing
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is going"))
	})

	//SUbrouting
	router.Route("/todo", func(r chi.Router) {
		//middleware
		r.Use(logMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Lista de todo"))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("adaug un todo"))
		})

	})

	fmt.Println("Server pornit pe http://localhost:3000")
	http.ListenAndServe(":3000", router)

	return
}
