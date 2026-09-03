-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS user_name FROM FEEDS
JOIN users ON FEEDS.user_id = users.id;


-- name: GetFeedByURL :one
SELECT feeds.name, feeds.id FROM feeds
WHERE url = $1;