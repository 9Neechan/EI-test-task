-- name: CreateService :one
INSERT INTO services (name, description, created_at)
VALUES ($1, $2, NOW())
RETURNING id;