-- name: CreateUser :one
INSERT INTO "user" (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)

RETURNING *;

-- name: GetUsers :many
SELECT * FROM "user";

-- name: GetUser :one
SELECT * FROM "user" WHERE name = $1;

-- name: Reset :exec
DELETE FROM "user";
