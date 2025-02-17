-- +goose Up
CREATE TABLE
	feed_follow (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		user_id UUID REFERENCES users (id) ON DELETE CASCADE,
		feed_id UUID REFERENCES feeds (id) ON DELETE CASCADE,
		CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id)
	);

-- +goose Down
DROP TABLE IF EXISTS feed_follow CASCADE;