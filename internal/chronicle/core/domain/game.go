package domain

import "github.com/google/uuid"

type Game struct {
	GameId uuid.UUID
	Name   string
	Type   string
	// TODO: Need to decide how many worlds to support in a game.. maybe just open it up?
	World *World
}
