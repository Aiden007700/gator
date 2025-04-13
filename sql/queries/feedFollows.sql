-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at, user_id, feed_id;

-- name: GetFeedFollowWithDetails :one
SELECT
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    feed_follows.user_id,
    feed_follows.feed_id,
    users.name AS user_name,
    feed.name AS feed_name
FROM feed_follows
         JOIN users ON feed_follows.user_id = users.id
         JOIN feed ON feed_follows.feed_id = feed.id
WHERE feed_follows.id = $1;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    feed_follows.user_id,
    feed_follows.feed_id,
    users.name AS user_name,
    feed.name AS feed_name
FROM feed_follows
         JOIN users ON feed_follows.user_id = users.id
         JOIN feed ON feed_follows.feed_id = feed.id
WHERE users.name = $1;

-- name: DeleteFeedFollowByUserAndUrl :exec
DELETE FROM feed_follows
    USING feed
WHERE feed_follows.feed_id = feed.id
  AND feed_follows.user_id = $1
  AND feed.url = $2;