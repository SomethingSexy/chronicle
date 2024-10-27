package domain

import "github.com/google/uuid"

// TODO: How should we handle invalid types here?
func NewGameCharacter(
	gameType GameType,
	id uuid.UUID,
	characterId uuid.UUID,
	characterType CharacterType,
	character map[string]interface{},
) GameCharacter {
	var gameCharacter GameCharacter
	if gameType == VTM {
		gameCharacter = NewVtmGameCharacter(id, characterId, characterType, character)
	}

	return gameCharacter
}

// This should probably be generic, need to
// test this against the schema compiler though.
// Maybe after it is valid, we can marshall to a strict type
// type GameCharacter[D Schema] struct {
// 	Data D
// 	Type GameType
// }

type GameCharacter interface {
	Validator
	GetType() CharacterType
	GetId() uuid.UUID
	GetCharacterId() uuid.UUID
}

func NewCharacterType(t string) CharacterType {
	switch t {
	case "npc":
		return NPC
	case "pc":
		return PC
	}

	return NPC
}

type CharacterType int

const (
	NPC = iota
	PC
)

func (t CharacterType) String() string {
	return [...]string{"npc", "pc"}[t]
}
