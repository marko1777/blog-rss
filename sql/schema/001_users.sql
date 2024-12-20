-- +goose Up
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE "user";
