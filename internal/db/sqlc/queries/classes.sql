-- name: CreateClass :one
INSERT INTO classes (id, name, section)
VALUES ($1, $2, $3)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetAnyClass :one
SELECT * FROM classes LIMIT 1;
