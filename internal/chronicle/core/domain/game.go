package domain

import (
	"github.com/google/uuid"
)

func NewGameType(t string) GameType {
	switch t {
	case "vtm":
		return VTM
	case "generic":
		return Generic
	}

	return Generic
}

// There will more than likely need to be a much more robust
// setup for different types of games, utilizing this for now.
//
// Game Type will be used to determine a lot of different things,
// including how characters are created
type GameType int

const (
	Generic = iota
	VTM
)

func (t GameType) String() string {
	return [...]string{"generic", "vtm"}[t]
}

type Game struct {
	GameId  uuid.UUID
	WorldId uuid.UUID
	Name    string
	Type    GameType
	World   World
}
