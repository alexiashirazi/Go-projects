package store

import (
	"context"
	"curs1_boilerplate/cmd/backend/model"
	"curs1_boilerplate/db"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

type dbUserStore struct {
	conn db.DBTX
}

func (d *dbUserStore) Add(u model.User) error {
	queries := db.New(d.conn)
	return queries.AddUser(context.Background(), db.AddUserParams{

		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	})
}

func (d *dbUserStore) GetByEmail(email string) (model.User, error) {
	queries := db.New(d.conn)
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func (d *dbUserStore) GetById(id string) (model.User, error) {
	queries := db.New(d.conn)

	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return model.User{}, err
	}

	user, err := queries.GetUserByID(context.Background(), uuid)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func (d *dbUserStore) GetAll() []model.User {
	queries := db.New(d.conn)
	dbUsers, err := queries.SelectUsers(context.Background())
	if err != nil {
		log.Println("error fetching users:", err)
		return nil
	}

	users := make([]model.User, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = model.User{
			ID:        u.ID.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  u.Password,
		}
	}
	return users
}

func (d *dbUserStore) Update(u model.User) error {
	queries := db.New(d.conn)
	return queries.UpdateUserPassword(context.Background(), db.UpdateUserPasswordParams{
		Email:    u.Email,
		Password: u.Password,
	})

}

func (d *dbUserStore) Delete(id string) error {
	queries := db.New(d.conn)

	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return err
	}

	return queries.DeleteUser(context.Background(), uuid)
}

func NewDbUserStore(conn db.DBTX) UserStore {
	return &dbUserStore{
		conn: conn,
	}
}
