package game

import (
	"os"
)

// TODO: I think this needs another sub property to work with the validator
type VtmGameCharacter struct {
	Name        string       `json:"name"`
	Disciplines []Discipline `json:"disciplines,omitempty"`
}

type Discipline struct {
	Name   string  `json:"name"`
	Level  int     `json:"level"`
	Powers []Power `json:"powers"`
}

type Power struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description,omitempty"`
}

func (g VtmGameCharacter) Schema() ([]byte, error) {
	return os.ReadFile("./vtm_v5_character_schema.json")
}
