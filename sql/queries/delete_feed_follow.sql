-- name: DeleteFeedFollow :execrows
DELETE FROM feed_follows WHERE id = $1;