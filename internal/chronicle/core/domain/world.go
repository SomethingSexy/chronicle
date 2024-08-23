package domain

import "github.com/google/uuid"

// This represents the game world.
// Right now there is 1:1 and this might just hold
// data specific to the game world
type World struct {
	GameId  uuid.UUID
	WorldId uuid.UUID
	Name    string
	// Root level locations
	Loocation []Location
}
