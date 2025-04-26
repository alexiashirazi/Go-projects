package store

import (
	"context"
	"curs1_boilerplate/cmd/exemplu_store/model"
	"curs1_boilerplate/db"
	"log"
)

type dbProductStore struct {
	conn db.DBTX
}

// Add implements ProductStore.
func (d *dbProductStore) Add(p model.Product) error {
	queries := db.New(d.conn)
	return queries.InsertProduct(context.Background(), db.InsertProductParams{
		Name:  p.Name,
		Value: int32(p.Value),
	})
}

// GetAll implements ProductStore.
func (d *dbProductStore) GetAll() []model.Product {
	queries := db.New(d.conn)
	dbProducts, err := queries.GetAllProducts(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	products := make([]model.Product, len(dbProducts))
	for i, dbP := range dbProducts {
		products[i] = model.Product{
			Name:  dbP.Name,
			Value: int(dbP.Value),
		}
	}

	return products
}

func NewDbProductStore(conn db.DBTX) ProductStore {
	return &dbProductStore{
		conn: conn,
	}
}
