package store

import (
	"context"
	"curs1_boilerplate/cmd/backend/model"
	"curs1_boilerplate/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type dbCategoryStore struct {
	conn db.DBTX
}

// AddCategory implements CategoryStore.
func (d *dbCategoryStore) AddCategory(c model.Category) error {
	queries := db.New(d.conn)
	return queries.InsertCategory(context.Background(), c.Name)
}

// DeleteCategory implements CategoryStore.
func (d *dbCategoryStore) DeleteCategory(id string) error {
	queries := db.New(d.conn)
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return err
	}
	return queries.DeleteCategory(context.Background(), uuid)
}

// GetAllCategories implements CategoryStore.
func (d *dbCategoryStore) GetAllCategories() ([]model.Category, error) {
	queries := db.New(d.conn)
	categories, err := queries.GetAllCategories(context.Background())
	if err != nil {
		return nil, err
	}
	var result []model.Category
	for _, c := range categories {
		result = append(result, model.Category{
			ID:   c.ID.String(),
			Name: c.Name,
		})
	}
	return result, nil
}

// GetCategoryByID implements CategoryStore.
func (d *dbCategoryStore) GetCategoryByID(id string) (model.Category, error) {
	queries := db.New(d.conn)
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return model.Category{}, err
	}
	category, err := queries.GetCategoryByID(context.Background(), uuid)
	if err != nil {
		return model.Category{}, err
	}
	return model.Category{
		ID:   category.ID.String(),
		Name: category.Name,
	}, nil
}

// UpdateCategory implements CategoryStore.
func (d *dbCategoryStore) UpdateCategory(c model.Category) error {
	queries := db.New(d.conn)

	var uuid pgtype.UUID
	if err := uuid.Scan(c.ID); err != nil {
		return err
	}

	return queries.UpdateCategory(context.Background(), db.UpdateCategoryParams{
		ID:   uuid,
		Name: c.Name,
	})
}

func NewDbCategoryStore(conn db.DBTX) CategoryStore {
	return &dbCategoryStore{conn: conn}
}
