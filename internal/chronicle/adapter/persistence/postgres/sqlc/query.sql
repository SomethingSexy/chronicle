-- name: GetGame :one
SELECT * FROM game
WHERE id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM game
ORDER BY name;

-- name: CreateGame :one
INSERT INTO game (
  name, type
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateGame :exec
UPDATE game
  set name = $2,
  type = $3
WHERE id = $1;

-- name: DeleteGame :exec
DELETE FROM game
WHERE id = $1;