-- +goose Up
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name TEXT NOT NULL,

    UNIQUE (name)
);

-- +goose Down
DROP TABLE users;
