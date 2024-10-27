package domain

import (
	"os"
	"path"

	"github.com/google/uuid"
)

func NewVtmGameCharacter(
	id uuid.UUID,
	characterId uuid.UUID,
	characterType CharacterType,
	character map[string]interface{},
) GameCharacter {
	return VtmGameCharacter{
		Id:            id,
		CharacterId:   characterId,
		Character:     character,
		CharacterType: characterType,
	}
}

// For now, keep it generic until we need something more
// conrete in code to deal with
type VtmGameCharacter struct {
	Id            uuid.UUID
	CharacterId   uuid.UUID
	Character     map[string]interface{}
	CharacterType CharacterType
}

func (g VtmGameCharacter) GetId() uuid.UUID {
	return g.Id
}

func (g VtmGameCharacter) GetCharacterId() uuid.UUID {
	return g.CharacterId
}

func (g VtmGameCharacter) GetSchema() ([]byte, error) {
	return os.ReadFile(path.Join("core", "domain", "schema", "vtm_v5_character_schema.json"))
}

func (g VtmGameCharacter) GetData() map[string]interface{} {
	return g.Character
}

func (g VtmGameCharacter) GetType() CharacterType {
	return g.CharacterType
}

// TODO: I think this needs another sub property to work with the validator
// type VtmGameCharacter struct {
// 	Name        string       `json:"name"`
// 	Disciplines []Discipline `json:"disciplines,omitempty"`
// }

// type Discipline struct {
// 	Name   string  `json:"name"`
// 	Level  int     `json:"level"`
// 	Powers []Power `json:"powers"`
// }

// type Power struct {
// 	Name        string `json:"name"`
// 	Level       int    `json:"level"`
// 	Description string `json:"description,omitempty"`
// }
