-- name: AddUser :exec
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4);

-- name: SelectUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2
WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
