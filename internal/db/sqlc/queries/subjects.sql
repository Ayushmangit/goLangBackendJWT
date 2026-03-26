-- name: CreateSubject :one
INSERT INTO subjects (id, name)
VALUES ($1, $2)
ON CONFLICT (name) DO NOTHING
RETURNING *;
