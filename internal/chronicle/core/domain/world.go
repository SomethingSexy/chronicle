package domain

import "github.com/google/uuid"

// This represents a generic world.
type World struct {
	WorldId uuid.UUID
	Name    string
	// Root level locations
	Locations  []Location
	Characters []Character
}
