-- name: AddContact :one
INSERT INTO contact_list( username,member_username,last_activity) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetContacts :many
SELECT * FROM contact_list
WHERE username=$1;