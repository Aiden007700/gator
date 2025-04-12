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