// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: posts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO
	posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)
ON CONFLICT (url) DO NOTHING
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.NullUUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostsFromUser = `-- name: GetPostsFromUser :many
SELECT
	posts.id, posts.created_at, posts.updated_at, title, url, description, published_at, posts.feed_id, ff.id, ff.created_at, ff.updated_at, user_id, ff.feed_id
FROM
	posts
JOIN feed_follow ff ON posts.feed_id = ff.feed_id
WHERE
	ff.user_id = $1	
ORDER BY
	published_at DESC
LIMIT
	$2
`

type GetPostsFromUserParams struct {
	UserID uuid.NullUUID
	Limit  int32
}

type GetPostsFromUserRow struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.NullUUID
	ID_2        uuid.UUID
	CreatedAt_2 time.Time
	UpdatedAt_2 time.Time
	UserID      uuid.NullUUID
	FeedID_2    uuid.NullUUID
}

func (q *Queries) GetPostsFromUser(ctx context.Context, arg GetPostsFromUserParams) ([]GetPostsFromUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsFromUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsFromUserRow
	for rows.Next() {
		var i GetPostsFromUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.UserID,
			&i.FeedID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
