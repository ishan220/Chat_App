-- name: CreateUser :one
INSERT INTO users (
    username,
    password
) VALUES ($1, $2) 
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 AND 
password= $2 LIMIT 1;

-- name: VerifyUser :one
SELECT * FROM users
WHERE username = $1
limit 1;