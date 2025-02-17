-- +goose Up
CREATE TABLE
	feeds (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP NOT NUll,
		updated_at TIMESTAMP NOT NULL,
		name TEXT UNIQUE NOT NULL,
		url TEXT UNIQUE NOT NULL,
		user_id UUID references users (id) ON DELETE CASCADE,
		last_fetched_at TIMESTAMP
	);

-- +goose Down
DROP TABLE IF EXISTS feeds CASCADE;