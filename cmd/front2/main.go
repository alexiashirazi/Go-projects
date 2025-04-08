package main

import (
	"curs1_boilerplate/cmd/front2/views/base"
	"curs1_boilerplate/cmd/front2/views/pages"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
wgo -file '\.templ$' templ generate &
wgo -file .go go run .
*/

func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		base.PageSkeleton(pages.Page1()).Render(r.Context(), w)
	})

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
