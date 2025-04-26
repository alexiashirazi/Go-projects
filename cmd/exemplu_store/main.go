package main

import (
	"context"
	"curs1_boilerplate/cmd/exemplu_store/model"
	"curs1_boilerplate/cmd/exemplu_store/store"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func AddProduct(s store.ProductStore, p model.Product) {
	s.Add(p)
}

func GetAllProducts(s store.ProductStore) []model.Product {
	return s.GetAll()
}

type Controller struct {
	store store.ProductStore
}

func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := c.store.GetAll()
	json.NewEncoder(w).Encode(products)
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	// s := store.NewMemoryStore()
	s := store.NewDbProductStore(conn)
	c := &Controller{store: s}

	r := chi.NewRouter()
	r.Get("/products", c.GetProducts)

	http.ListenAndServe("localhost:3000", r)

	// AddProduct(s, model.Product{
	// 	Name:  "ceapa",
	// 	Value: 999,
	// })
	// AddProduct(s, model.Product{
	// 	Name:  "usturoy",
	// 	Value: 1000,
	// })
	// AddProduct(s, model.Product{
	// 	Name:  "paine",
	// 	Value: 50,
	// })
	//
	// for _, p := range s.GetAll() {
	// 	fmt.Println(p)
	// }
}
