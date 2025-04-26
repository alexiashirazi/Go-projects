package store

import (
	"curs1_boilerplate/cmd/backend/model"
)

type ProductStore interface {
	GetAll() ([]model.Product, error)
	Add(p model.Product) error
	GetByID(id string) (model.Product, error)
	Update(p model.Product) error
	Delete(id string) error
	GetByCategory(categoryID string) ([]model.Product, error)
}
