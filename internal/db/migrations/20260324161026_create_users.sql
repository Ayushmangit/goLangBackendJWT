-- +goose Up
CREATE TABLE users(
		id uuid PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;
