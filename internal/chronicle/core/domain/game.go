package domain

import "github.com/google/uuid"

type Game struct {
	GameId uuid.UUID
	Name   string
	Type   string
	Worlds []World
}
