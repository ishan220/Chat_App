// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: contact_list.sql

package db

import (
	"context"
)

const addContact = `-- name: AddContact :one
INSERT INTO contact_list( username,member_username,last_activity) 
VALUES ($1, $2, $3) 
RETURNING username, member_username, last_activity
`

type AddContactParams struct {
	Username       string `json:"username"`
	MemberUsername string `json:"member_username"`
	LastActivity   int64  `json:"last_activity"`
}

func (q *Queries) AddContact(ctx context.Context, arg AddContactParams) (ContactList, error) {
	row := q.db.QueryRowContext(ctx, addContact, arg.Username, arg.MemberUsername, arg.LastActivity)
	var i ContactList
	err := row.Scan(&i.Username, &i.MemberUsername, &i.LastActivity)
	return i, err
}

const getContacts = `-- name: GetContacts :many
SELECT username, member_username, last_activity FROM contact_list
WHERE username=$1
`

func (q *Queries) GetContacts(ctx context.Context, username string) ([]ContactList, error) {
	rows, err := q.db.QueryContext(ctx, getContacts, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ContactList
	for rows.Next() {
		var i ContactList
		if err := rows.Scan(&i.Username, &i.MemberUsername, &i.LastActivity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
