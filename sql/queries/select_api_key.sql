-- name: SelectByApiKey :one
SELECT * from users
WHERE api_key = $1;