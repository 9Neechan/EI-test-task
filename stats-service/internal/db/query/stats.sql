-- name: PostCall :one
INSERT INTO stats (user_id, service_id, count, created_at)
VALUES ($1, $2, 1, NOW())
ON CONFLICT (user_id, service_id) 
DO UPDATE SET count = stats.count + 1
RETURNING *;

-- name: GetStats :many
SELECT s.user_id, s.service_id, s.count, u.name AS user_name, srv.name AS service_name
FROM stats s
JOIN users u ON s.user_id = u.id
JOIN services srv ON s.service_id = srv.id
WHERE (s.user_id = $1 OR $1 = 0)
AND (s.service_id = $2 OR $2 = 0)
ORDER BY s.count DESC
LIMIT $3 OFFSET $4;