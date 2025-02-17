-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follow (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT 
    inserted.id, 
    inserted.created_at, 
    inserted.updated_at, 
    users.name AS user_name, 
    feeds.name AS feed_name
FROM inserted
JOIN users ON users.id = inserted.user_id
JOIN feeds ON feeds.id = inserted.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT 
    ff.id, 
    ff.created_at, 
    ff.updated_at, 
    u.name AS user_name, 
    f.name AS feed_name
FROM feed_follow ff
JOIN users u ON u.id = ff.user_id
JOIN feeds f ON f.id = ff.feed_id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow
WHERE user_id = (SELECT id FROM users u WHERE u.name = $1)
AND feed_id = (SELECT id FROM feeds f WHERE f.url = $2);