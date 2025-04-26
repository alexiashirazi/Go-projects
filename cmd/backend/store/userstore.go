package store

import (
	"curs1_boilerplate/cmd/backend/model"
)

type UserStore interface {
	GetAll() []model.User
	Add(u model.User) error
	GetByEmail(email string) (model.User, error)
	GetById(id string) (model.User, error)
	Update(u model.User) error
	Delete(id string) error
}
