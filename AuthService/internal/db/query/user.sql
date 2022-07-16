-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = ?
LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (username, passwd, phone)
VALUES (?, ?, ?);