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
  game_id, name, type, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (game_id) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  updated_at = EXCLUDED.updated_at
RETURNING *;

-- name: UpdateGame :exec
UPDATE game
  set name = $2,
  type = $3,
  updated_at = $4
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
  world_id, game_id, name, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (world_id) DO UPDATE SET
  name = EXCLUDED.name,
  updated_at = EXCLUDED.updated_at
RETURNING *;

-- name: UpdateWorld :exec
UPDATE world
  set name = $2
WHERE world_id = $1;

-- name: DeleteWorld :exec
DELETE FROM world
WHERE world_id = $1;

-- name: CreateLocation :one
INSERT INTO location (
  location_id, world_id, game_id, type, name, path, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (location_id) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  path = EXCLUDED.path,
  updated_at = EXCLUDED.updated_at
RETURNING *;

-- name: GetWorldLocations :many
SELECT * FROM location
JOIN world ON location.world_id = world.id
JOIN game ON location.game_id = game.id
WHERE game.game_id = $1 and
world.world_id = $2;

-- name: CreateCharacter :one
INSERT INTO character (
  character_id, name, description, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (character_id) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  updated_at = EXCLUDED.updated_at
RETURNING *;

-- name: GetCharacter :one
SELECT * FROM character
WHERE id = $1 LIMIT 1;

-- name: GetCharacterFromUuid :one
SELECT * FROM character
WHERE character.character_id = $1 LIMIT 1;

-- name: UpsertCharacterToGameWorld :exec
INSERT INTO world_character (
  world_character_id, world_id, character_id, character_type, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
ON CONFLICT (world_id, character_id) DO UPDATE SET
  character_type = EXCLUDED.character_type,
  updated_at = EXCLUDED.updated_at;

-- name: GetWorldCharacters :many
SELECT * FROM character
JOIN world ON world.world_id = $1
JOIN world_character ON world_character.world_id = world.id;
