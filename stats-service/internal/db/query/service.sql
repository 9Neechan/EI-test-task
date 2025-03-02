-- name: CreateService :one
INSERT INTO services (name, description, price, created_at)
VALUES ($1, $2, $3, NOW())
RETURNING *;