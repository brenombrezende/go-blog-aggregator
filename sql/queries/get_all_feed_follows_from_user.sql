-- name: SelectAllFeedFollows :many
SELECT * from feed_follows
WHERE user_id = $1;