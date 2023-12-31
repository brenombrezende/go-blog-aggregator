// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: get_user_by_api_key.sql

package database

import (
	"context"
)

const selectByApiKey = `-- name: SelectByApiKey :one
SELECT id, created_at, updated_at, name, api_key from users
WHERE api_key = $1
`

func (q *Queries) SelectByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, selectByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
