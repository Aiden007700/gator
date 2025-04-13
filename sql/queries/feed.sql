-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, user_id)
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
SELECT * FROM feed;

-- name: GetFeedsWithUser :many
SELECT * FROM feed JOIN users ON feed.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT * FROM feed where feed.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feed
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feed
ORDER BY last_fetched_at NULLS FIRST
    LIMIT 1;