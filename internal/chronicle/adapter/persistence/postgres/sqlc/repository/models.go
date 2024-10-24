// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Character struct {
	ID          int64
	CharacterID uuid.UUID
	Name        string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type Game struct {
	ID        int64
	GameID    uuid.UUID
	WorldID   int64
	Name      string
	Type      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type GameCharacter struct {
	ID              int64
	GameCharacterID uuid.UUID
	GameID          int64
	CharacterID     int64
	CharacterType   string
	Character       []byte
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
}

type Location struct {
	ID         int64
	LocationID uuid.UUID
	WorldID    int64
	Type       string
	Name       string
	Path       pgtype.Text
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type World struct {
	ID        int64
	WorldID   uuid.UUID
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type WorldCharacter struct {
	ID               int64
	WorldCharacterID uuid.UUID
	CharacterID      int64
	WorldID          int64
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
}
