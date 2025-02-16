-- +goose Up
CREATE TABLE
	users (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP NOT NUll,
		updated_at TIMESTAMP NOT NULL,
		name TEXT UNIQUE NOT NULL
	);

-- +goose Down
DROP TABLE IF EXISTS users CASCADE;