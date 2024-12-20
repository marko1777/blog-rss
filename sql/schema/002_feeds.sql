-- +goose Up
CREATE TABLE IF NOT EXISTS "feed"(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "feed";
