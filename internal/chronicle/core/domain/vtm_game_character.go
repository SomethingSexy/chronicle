package domain

import (
	"os"
	"path"
)

func NewVtmGameCharacter(characterType CharacterType, character map[string]interface{}) GameCharacter {
	return VtmGameCharacter{
		Character:     character,
		CharacterType: characterType,
	}
}

// For now, keep it generic until we need something more
// conrete in code to deal with
type VtmGameCharacter struct {
	Character     map[string]interface{}
	CharacterType CharacterType
}

func (g VtmGameCharacter) Schema() ([]byte, error) {
	return os.ReadFile(path.Join("core", "domain", "schema", "vtm_v5_character_schema.json"))
}

func (g VtmGameCharacter) Data() any {
	return g.Character
}

func (g VtmGameCharacter) Type() CharacterType {
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
