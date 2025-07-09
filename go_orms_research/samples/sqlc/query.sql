-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (id, name)
VALUES (?, ?);
