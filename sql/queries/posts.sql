-- name: CreatePost :one
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
RETURNING *;

-- name: GetPostsFromUser :many
SELECT
	*
FROM
	posts
JOIN feed_follow ff ON posts.feed_id = ff.feed_id
WHERE
	ff.user_id = $1	
ORDER BY
	published_at DESC
LIMIT
	$2;