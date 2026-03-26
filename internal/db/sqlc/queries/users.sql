-- name: CreateUser :one
INSERT INTO users (id, email, password, role)
VALUES ($1, $2, $3, $4)
ON CONFLICT (email) DO NOTHING
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: FindByID :one
SELECT * FROM users where id = $1;
