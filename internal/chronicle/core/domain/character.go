package domain

import "github.com/google/uuid"

// At its core this struct will
// represent a basic character that is independent
// of game or world.
// This might allow us to great detailed characters
// upfront but then attach them to games
// Once attached to a game, more information can be
// added to the character
type Character struct {
	CharacterId uuid.UUID
	Name        string
	Description string
}
