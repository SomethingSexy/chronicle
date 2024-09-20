package domain

import "github.com/google/uuid"

// Represents a location within a world
// this can be anything
type Location struct {
	LocationId uuid.UUID
	WorldId    uuid.UUID
	Type       string
	Name       string
	Path       []uuid.UUID
}
