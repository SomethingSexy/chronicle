// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addCharacterToGameWorld = `-- name: AddCharacterToGameWorld :exec
INSERT INTO world_character (
  world_character_id, world_id, character_id, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (world_id, character_id) DO UPDATE SET
  updated_at = EXCLUDED.updated_at
`

type AddCharacterToGameWorldParams struct {
	WorldCharacterID uuid.UUID
	WorldID          int64
	CharacterID      int64
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
}

func (q *Queries) AddCharacterToGameWorld(ctx context.Context, arg AddCharacterToGameWorldParams) error {
	_, err := q.db.Exec(ctx, addCharacterToGameWorld,
		arg.WorldCharacterID,
		arg.WorldID,
		arg.CharacterID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createCharacter = `-- name: CreateCharacter :one
INSERT INTO character (
  character_id, name, description, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (character_id) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  updated_at = EXCLUDED.updated_at
RETURNING id, character_id, name, description, created_at, updated_at
`

type CreateCharacterParams struct {
	CharacterID uuid.UUID
	Name        string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error) {
	row := q.db.QueryRow(ctx, createCharacter,
		arg.CharacterID,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.CharacterID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createGame = `-- name: CreateGame :one
INSERT INTO game (
  game_id, name, type, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (game_id) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  updated_at = EXCLUDED.updated_at
RETURNING id, game_id, name, type, created_at, updated_at
`

type CreateGameParams struct {
	GameID    uuid.UUID
	Name      string
	Type      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRow(ctx, createGame,
		arg.GameID,
		arg.Name,
		arg.Type,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Name,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createLocation = `-- name: CreateLocation :one
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
RETURNING id, location_id, game_id, world_id, type, name, path, created_at, updated_at
`

type CreateLocationParams struct {
	LocationID uuid.UUID
	WorldID    int64
	GameID     int64
	Type       string
	Name       string
	Path       pgtype.Text
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRow(ctx, createLocation,
		arg.LocationID,
		arg.WorldID,
		arg.GameID,
		arg.Type,
		arg.Name,
		arg.Path,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.LocationID,
		&i.GameID,
		&i.WorldID,
		&i.Type,
		&i.Name,
		&i.Path,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWorld = `-- name: CreateWorld :one
INSERT INTO world (
  world_id, game_id, name, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (world_id) DO UPDATE SET
  name = EXCLUDED.name,
  updated_at = EXCLUDED.updated_at
RETURNING id, world_id, game_id, name, created_at, updated_at
`

type CreateWorldParams struct {
	WorldID   uuid.UUID
	GameID    int64
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) CreateWorld(ctx context.Context, arg CreateWorldParams) (World, error) {
	row := q.db.QueryRow(ctx, createWorld,
		arg.WorldID,
		arg.GameID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i World
	err := row.Scan(
		&i.ID,
		&i.WorldID,
		&i.GameID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteGame = `-- name: DeleteGame :exec
DELETE FROM game
WHERE game_id = $1
`

func (q *Queries) DeleteGame(ctx context.Context, gameID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteGame, gameID)
	return err
}

const deleteWorld = `-- name: DeleteWorld :exec
DELETE FROM world
WHERE world_id = $1
`

func (q *Queries) DeleteWorld(ctx context.Context, worldID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteWorld, worldID)
	return err
}

const getCharacter = `-- name: GetCharacter :one
SELECT id, character_id, name, description, created_at, updated_at FROM character
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCharacter(ctx context.Context, id int64) (Character, error) {
	row := q.db.QueryRow(ctx, getCharacter, id)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.CharacterID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCharacterFromUuid = `-- name: GetCharacterFromUuid :one
SELECT id, character_id, name, description, created_at, updated_at FROM character
WHERE character.character_id = $1 LIMIT 1
`

func (q *Queries) GetCharacterFromUuid(ctx context.Context, characterID uuid.UUID) (Character, error) {
	row := q.db.QueryRow(ctx, getCharacterFromUuid, characterID)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.CharacterID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGame = `-- name: GetGame :one
SELECT id, game_id, name, type, created_at, updated_at FROM game
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGame(ctx context.Context, id int64) (Game, error) {
	row := q.db.QueryRow(ctx, getGame, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Name,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGameFromUuid = `-- name: GetGameFromUuid :one
SELECT id, game_id, name, type, created_at, updated_at FROM game
WHERE game.game_id = $1 LIMIT 1
`

func (q *Queries) GetGameFromUuid(ctx context.Context, gameID uuid.UUID) (Game, error) {
	row := q.db.QueryRow(ctx, getGameFromUuid, gameID)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Name,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGameWorlds = `-- name: GetGameWorlds :many
SELECT world.id, world_id, world.game_id, world.name, world.created_at, world.updated_at, game.id, game.game_id, game.name, type, game.created_at, game.updated_at FROM world
JOIN game ON world.game_id = game.id
WHERE game.game_id = $1
`

type GetGameWorldsRow struct {
	ID          int64
	WorldID     uuid.UUID
	GameID      int64
	Name        string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	ID_2        int64
	GameID_2    uuid.UUID
	Name_2      string
	Type        string
	CreatedAt_2 pgtype.Timestamptz
	UpdatedAt_2 pgtype.Timestamptz
}

func (q *Queries) GetGameWorlds(ctx context.Context, gameID uuid.UUID) ([]GetGameWorldsRow, error) {
	rows, err := q.db.Query(ctx, getGameWorlds, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGameWorldsRow
	for rows.Next() {
		var i GetGameWorldsRow
		if err := rows.Scan(
			&i.ID,
			&i.WorldID,
			&i.GameID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.GameID_2,
			&i.Name_2,
			&i.Type,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorld = `-- name: GetWorld :one
SELECT id, world_id, game_id, name, created_at, updated_at FROM world
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWorld(ctx context.Context, id int64) (World, error) {
	row := q.db.QueryRow(ctx, getWorld, id)
	var i World
	err := row.Scan(
		&i.ID,
		&i.WorldID,
		&i.GameID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getWorldCharacters = `-- name: GetWorldCharacters :many
SELECT character.id, character.character_id, character.name, description, character.created_at, character.updated_at, world.id, world.world_id, game_id, world.name, world.created_at, world.updated_at, world_character.id, world_character_id, world_character.character_id, world_character.world_id, world_character.created_at, world_character.updated_at FROM character
JOIN world ON world.world_id = $1
JOIN world_character ON world_character.world_id = world.id
`

type GetWorldCharactersRow struct {
	ID               int64
	CharacterID      uuid.UUID
	Name             string
	Description      pgtype.Text
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
	ID_2             int64
	WorldID          uuid.UUID
	GameID           int64
	Name_2           string
	CreatedAt_2      pgtype.Timestamptz
	UpdatedAt_2      pgtype.Timestamptz
	ID_3             int64
	WorldCharacterID uuid.UUID
	CharacterID_2    int64
	WorldID_2        int64
	CreatedAt_3      pgtype.Timestamptz
	UpdatedAt_3      pgtype.Timestamptz
}

func (q *Queries) GetWorldCharacters(ctx context.Context, worldID uuid.UUID) ([]GetWorldCharactersRow, error) {
	rows, err := q.db.Query(ctx, getWorldCharacters, worldID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWorldCharactersRow
	for rows.Next() {
		var i GetWorldCharactersRow
		if err := rows.Scan(
			&i.ID,
			&i.CharacterID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.WorldID,
			&i.GameID,
			&i.Name_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.ID_3,
			&i.WorldCharacterID,
			&i.CharacterID_2,
			&i.WorldID_2,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorldFromUuid = `-- name: GetWorldFromUuid :one
SELECT id, world_id, game_id, name, created_at, updated_at FROM world
WHERE world_id = $1 LIMIT 1
`

func (q *Queries) GetWorldFromUuid(ctx context.Context, worldID uuid.UUID) (World, error) {
	row := q.db.QueryRow(ctx, getWorldFromUuid, worldID)
	var i World
	err := row.Scan(
		&i.ID,
		&i.WorldID,
		&i.GameID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getWorldLocations = `-- name: GetWorldLocations :many
SELECT location.id, location_id, location.game_id, location.world_id, location.type, location.name, path, location.created_at, location.updated_at, world.id, world.world_id, world.game_id, world.name, world.created_at, world.updated_at, game.id, game.game_id, game.name, game.type, game.created_at, game.updated_at FROM location
JOIN world ON location.world_id = world.id
JOIN game ON location.game_id = game.id
WHERE game.game_id = $1 and
world.world_id = $2
`

type GetWorldLocationsParams struct {
	GameID  uuid.UUID
	WorldID uuid.UUID
}

type GetWorldLocationsRow struct {
	ID          int64
	LocationID  uuid.UUID
	GameID      int64
	WorldID     int64
	Type        string
	Name        string
	Path        pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	ID_2        int64
	WorldID_2   uuid.UUID
	GameID_2    int64
	Name_2      string
	CreatedAt_2 pgtype.Timestamptz
	UpdatedAt_2 pgtype.Timestamptz
	ID_3        int64
	GameID_3    uuid.UUID
	Name_3      string
	Type_2      string
	CreatedAt_3 pgtype.Timestamptz
	UpdatedAt_3 pgtype.Timestamptz
}

func (q *Queries) GetWorldLocations(ctx context.Context, arg GetWorldLocationsParams) ([]GetWorldLocationsRow, error) {
	rows, err := q.db.Query(ctx, getWorldLocations, arg.GameID, arg.WorldID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWorldLocationsRow
	for rows.Next() {
		var i GetWorldLocationsRow
		if err := rows.Scan(
			&i.ID,
			&i.LocationID,
			&i.GameID,
			&i.WorldID,
			&i.Type,
			&i.Name,
			&i.Path,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.WorldID_2,
			&i.GameID_2,
			&i.Name_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.ID_3,
			&i.GameID_3,
			&i.Name_3,
			&i.Type_2,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGames = `-- name: ListGames :many
SELECT id, game_id, name, type, created_at, updated_at FROM game
ORDER BY name
`

func (q *Queries) ListGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.Query(ctx, listGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.Name,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listWorlds = `-- name: ListWorlds :many
SELECT id, world_id, game_id, name, created_at, updated_at FROM world
ORDER BY name
`

func (q *Queries) ListWorlds(ctx context.Context) ([]World, error) {
	rows, err := q.db.Query(ctx, listWorlds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []World
	for rows.Next() {
		var i World
		if err := rows.Scan(
			&i.ID,
			&i.WorldID,
			&i.GameID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGame = `-- name: UpdateGame :exec
UPDATE game
  set name = $2,
  type = $3,
  updated_at = $4
WHERE game_id = $1
`

type UpdateGameParams struct {
	GameID    uuid.UUID
	Name      string
	Type      string
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) error {
	_, err := q.db.Exec(ctx, updateGame,
		arg.GameID,
		arg.Name,
		arg.Type,
		arg.UpdatedAt,
	)
	return err
}

const updateWorld = `-- name: UpdateWorld :exec
UPDATE world
  set name = $2
WHERE world_id = $1
`

type UpdateWorldParams struct {
	WorldID uuid.UUID
	Name    string
}

func (q *Queries) UpdateWorld(ctx context.Context, arg UpdateWorldParams) error {
	_, err := q.db.Exec(ctx, updateWorld, arg.WorldID, arg.Name)
	return err
}
