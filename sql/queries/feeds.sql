-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
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
SELECT * FROM feeds;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: GetFeedById :one
SELECT * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetFeedsWithUserName :many
SELECT feeds.*, users.name as username FROM feeds
LEFT JOIN users ON feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
