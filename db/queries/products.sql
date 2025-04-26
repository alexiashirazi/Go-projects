-- name: InsertProduct :exec
INSERT INTO products (
    category_id,
    device_type,
    model,
    color,
    storage,
    battery_health,
    processor,
    ram,
    description,
    created_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);


SELECT * FROM products;

-- name: GetAllProducts :many
SELECT * FROM products ;

-- name: GetProductsByCategory :many
SELECT * FROM products
WHERE category_id = $1;

-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET 
    category_id = $2,
    device_type = $3,
    model = $4,
    color = $5,
    storage = $6,
    battery_health = $7,
    processor = $8,
    ram = $9,
    description = $10
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;



-- name: InsertCategory :exec
INSERT INTO categories (name) VALUES ($1);

-- name: GetAllCategories :many
SELECT id, name FROM categories;

-- name: GetCategoryByID :one
SELECT id, name FROM categories
WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories
SET name = $2
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;