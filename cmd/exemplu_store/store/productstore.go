package store

import "curs1_boilerplate/cmd/exemplu_store/model"

type ProductStore interface {
	GetAll() []model.Product
	Add(p model.Product) error
}

