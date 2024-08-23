-- name: GetGame :one
SELECT * FROM game
WHERE id = $1 LIMIT 1;

-- name: GetGameFromUuid :one
SELECT * FROM game
WHERE game.game_id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM game
ORDER BY name;

-- name: CreateGame :one
INSERT INTO game (
  game_id, name, type
) VALUES (
  $1, $2, $3
)
ON CONFLICT (game_id) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type
RETURNING *;

-- name: UpdateGame :exec
UPDATE game
  set name = $2,
  type = $3
WHERE game_id = $1;

-- name: DeleteGame :exec
DELETE FROM game
WHERE game_id = $1;

-- name: GetWorld :one
SELECT * FROM world
WHERE id = $1 LIMIT 1;

-- name: GetWorldFromUuid :one
SELECT * FROM world
WHERE world_id = $1 LIMIT 1;

-- name: GetGameWorlds :many
SELECT * FROM world
JOIN game ON world.game_id = game.id
WHERE game.game_id = $1;

-- name: ListWorlds :many
SELECT * FROM world
ORDER BY name;

-- name: CreateWorld :one
INSERT INTO world (
  world_id, game_id, name
) VALUES (
  $1, $2, $3
)
ON CONFLICT (world_id) DO UPDATE SET
  name = EXCLUDED.name
RETURNING *;

-- name: UpdateWorld :exec
UPDATE world
  set name = $2
WHERE world_id = $1;

-- name: DeleteWorld :exec
DELETE FROM world
WHERE world_id = $1;