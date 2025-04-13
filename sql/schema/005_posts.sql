-- +goose Up
CREATE TABLE posts (
                       id UUID PRIMARY KEY,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       title TEXT NOT NULL,
                       url TEXT UNIQUE NOT NULL,
                       description TEXT,
                       published_at TIMESTAMP WITH TIME ZONE,
                       feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;