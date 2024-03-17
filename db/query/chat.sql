-- name: CreateChat :one
INSERT INTO chat( id,from_user,to_user,message,created_at) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetChatHistory :many
SELECT * FROM chat 
WHERE (from_user=$1  OR to_user=$1) 
AND (from_user=$2 OR to_user=$2);