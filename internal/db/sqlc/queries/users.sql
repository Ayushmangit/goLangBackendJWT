-- name: CreateUser :one
INSERT INTO users (id,email,password)
VALUES ($1,$2,$3)
returning *;

-- name: GetUserByEmail :one
SELECT * FROM users where email = $1;

-- name: FindByID :one
SELECT * FROM users where id = $1;
