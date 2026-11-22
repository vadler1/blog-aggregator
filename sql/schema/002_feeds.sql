-- +goose Up
CREATE TABLE feeds (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    UNIQUE (user_id, url)
);

-- +goose Down
DROP TABLE feeds;
