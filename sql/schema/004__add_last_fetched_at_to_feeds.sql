-- +goose Up
ALTER TABLE feed
    ADD COLUMN last_fetched_at TIMESTAMP WITH TIME ZONE NULL;

-- +goose Down
ALTER TABLE feed
DROP COLUMN last_fetched_at;