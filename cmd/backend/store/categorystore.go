package store

import "curs1_boilerplate/cmd/backend/model"

type CategoryStore interface {
	GetAllCategories() ([]model.Category, error)
	AddCategory(c model.Category) error
	GetCategoryByID(id string) (model.Category, error)
	UpdateCategory(c model.Category) error
	DeleteCategory(id string) error
}
