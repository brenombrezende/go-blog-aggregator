// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: get_all_feed_follows_from_user.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const selectAllFeedFollows = `-- name: SelectAllFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id from feed_follows
WHERE user_id = $1
`

func (q *Queries) SelectAllFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, selectAllFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
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
