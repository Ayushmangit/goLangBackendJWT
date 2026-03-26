-- name: CreateStudent :one
INSERT INTO students (id, user_id, full_name, roll_number, class_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (roll_number) DO NOTHING
RETURNING *;
