package store

import (
	"context"
	"curs1_boilerplate/cmd/backend/model"
	"curs1_boilerplate/db"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type dbProductStore struct {
	conn db.DBTX
}

// GetByUser implements ProductStore.
func (d *dbProductStore) GetByUser(userID string) ([]model.Product, error) {
	queries := db.New(d.conn)

	var uuid pgtype.UUID
	if err := uuid.Scan(userID); err != nil {
		return nil, err
	}

	products, err := queries.GetProductsByUser(context.Background(), uuid)
	if err != nil {
		return nil, err
	}

	var result []model.Product
	for _, p := range products {
		result = append(result, model.Product{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			CategoryID:    p.CategoryID.String(),
			DeviceType:    p.DeviceType,
			Model:         p.Model,
			Color:         p.Color.String,
			Storage:       p.Storage.String,
			BatteryHealth: p.BatteryHealth.String,
			Processor:     p.Processor.String,
			Ram:           p.Ram.String,
			Description:   p.Description.String,
			CreatedAt:     p.CreatedAt.Time.String(),
		})
	}
	return result, nil
}

// Add implements ProductStore.

func (d *dbProductStore) Add(p model.Product) error {
	queries := db.New(d.conn)
	fmt.Println("Full product data:", p)

	// Convert string CategoryID → pgtype.UUID
	var categoryUUID pgtype.UUID
	if err := categoryUUID.Scan(p.CategoryID); err != nil {
		fmt.Println("CategoryID conversion error:", err)
		return fmt.Errorf("CategoryID conversion error: %w", err)
	}

	var userUUID pgtype.UUID
	if err := userUUID.Scan(p.UserID); err != nil {
		fmt.Println("UserID conversion error:", err)
		return fmt.Errorf("UserID conversion error: %w", err)
	}

	// Convert string CreatedAt → time.Time → pgtype.Timestamp
	var createdAt pgtype.Timestamp
	if p.CreatedAt != "" {
		parsedCreatedAt, err := time.Parse(time.RFC3339, p.CreatedAt)
		if err != nil {
			fmt.Println("CreatedAt parsing error:", err)
			return fmt.Errorf("CreatedAt parsing error: %w", err)
		}
		if err := createdAt.Scan(parsedCreatedAt); err != nil {
			fmt.Println("CreatedAt scan error:", err)
			return fmt.Errorf("CreatedAt scan error: %w", err)
		}
	} else {
		// fallback dacă nu trimite frontend data
		createdAt.Time = time.Now()
		createdAt.Valid = true
	}

	// // Check for required fields
	// if p.DeviceType == "" {
	// 	return fmt.Errorf("DeviceType is required")
	// }
	// if p.Model == "" {
	// 	return fmt.Errorf("Model is required")
	// }

	// Create parameters for database insertion
	params := db.InsertProductParams{
		UserID:        userUUID,
		CategoryID:    categoryUUID,
		DeviceType:    p.DeviceType,
		Model:         p.Model,
		Color:         pgtype.Text{String: p.Color, Valid: p.Color != ""},
		Storage:       pgtype.Text{String: p.Storage, Valid: p.Storage != ""},
		BatteryHealth: pgtype.Text{String: p.BatteryHealth, Valid: p.BatteryHealth != ""},
		Processor:     pgtype.Text{String: p.Processor, Valid: p.Processor != ""},
		Ram:           pgtype.Text{String: p.Ram, Valid: p.Ram != ""},
		Description:   pgtype.Text{String: p.Description, Valid: p.Description != ""},
		CreatedAt:     createdAt,
	}

	// Print the parameters being sent to the database
	fmt.Printf("Inserting with params: %+v\n", params)

	// Attempt the insert
	err := queries.InsertProduct(context.Background(), params)
	if err != nil {
		fmt.Println("Database insertion error:", err)

	}

	return nil
}

// Delete implements ProductStore.
func (d *dbProductStore) Delete(id string) error {
	queries := db.New(d.conn)

	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return err
	}

	return queries.DeleteProduct(context.Background(), uuid)
}

// GetAll implements ProductStore.
func (d *dbProductStore) GetAll() ([]model.Product, error) {
	queries := db.New(d.conn)
	products, err := queries.GetAllProducts(context.Background())
	if err != nil {
		return nil, err
	}

	var result []model.Product
	for _, p := range products {
		result = append(result, model.Product{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			CategoryID:    p.CategoryID.String(),
			DeviceType:    p.DeviceType,
			Model:         p.Model,
			Color:         p.Color.String,
			Storage:       p.Storage.String,
			BatteryHealth: p.BatteryHealth.String,
			Processor:     p.Processor.String,
			Ram:           p.Ram.String,
			Description:   p.Description.String,
			CreatedAt:     p.CreatedAt.Time.String(),
		})
	}
	return result, nil
}

// GetByCategory implements ProductStore.
func (d *dbProductStore) GetByCategory(categoryID string) ([]model.Product, error) {
	queries := db.New(d.conn)
	var uuid pgtype.UUID
	if err := uuid.Scan(categoryID); err != nil {
		return []model.Product{}, err
	}

	product, err := queries.GetProductsByCategory(context.Background(), uuid)
	if err != nil {
		return []model.Product{}, err
	}
	var result []model.Product

	for _, p := range product {
		result = append(result, model.Product{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			CategoryID:    p.CategoryID.String(),
			DeviceType:    p.DeviceType,
			Model:         p.Model,
			Color:         p.Color.String,
			Storage:       p.Storage.String,
			BatteryHealth: p.BatteryHealth.String,
			Processor:     p.Processor.String,
			Ram:           p.Ram.String,
			Description:   p.Description.String,
			CreatedAt:     p.CreatedAt.Time.String(),
		})
	}

	return result, nil
}

// GetByID implements ProductStore.
func (d *dbProductStore) GetByID(id string) (model.Product, error) {
	queries := db.New(d.conn)
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return model.Product{}, err
	}

	p, err := queries.GetProductByID(context.Background(), uuid)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:            p.ID.String(),
		UserID:        p.UserID.String(),
		CategoryID:    p.CategoryID.String(),
		DeviceType:    p.DeviceType,
		Model:         p.Model,
		Color:         p.Color.String,
		Storage:       p.Storage.String,
		BatteryHealth: p.BatteryHealth.String,
		Processor:     p.Processor.String,
		Ram:           p.Ram.String,
		Description:   p.Description.String,
		CreatedAt:     p.CreatedAt.Time.String(),
	}, nil

}

// Update implements ProductStore.
func (d *dbProductStore) Update(p model.Product) error {
	queries := db.New(d.conn)
	return queries.UpdateProduct(context.Background(), db.UpdateProductParams{
		DeviceType:    p.DeviceType,
		Model:         p.Model,
		Color:         pgtype.Text{String: p.Color, Valid: p.Color != ""},
		Storage:       pgtype.Text{String: p.Storage, Valid: p.Storage != ""},
		BatteryHealth: pgtype.Text{String: p.BatteryHealth, Valid: p.BatteryHealth != ""},
		Processor:     pgtype.Text{String: p.Processor, Valid: p.Processor != ""},
		Ram:           pgtype.Text{String: p.Ram, Valid: p.Ram != ""},
		Description:   pgtype.Text{String: p.Description, Valid: p.Description != ""},
	})
}

func NewDbProductStore(conn db.DBTX) ProductStore {
	return &dbProductStore{
		conn: conn,
	}
}
