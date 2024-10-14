package domain

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
	"github.com/kaptinlin/jsonschema"
)

// This should probably be generic, need to
// test this against the schema compiler though.
// Maybe after it is valid, we can marshall to a strict type
type GameCharacter[D interface{}] struct {
	Data D
}

func (g GameCharacter[D]) Validate() (bool, error) {
	// For now just read on load (for testing purposes), we can deal with caching this later
	var vtmV5CharacterSchema, err = os.ReadFile("./game/vtm_v5_character_schema.json")
	if err != nil {
		return false, err
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(vtmV5CharacterSchema)
	if err != nil {
		return false, err
	}

	// Need to figure this out but I believe the validator
	// requires it to be a map of interfaces.
	// For now we are going to marshall and unmarshall back into
	// a generic interface of maps so we can validate the schema
	// Hide it all here
	dataAsByte, err := json.Marshal(g.Data)
	if err != nil {
		return false, err
	}

	var dataAsInterface interface{}
	err = json.Unmarshal(dataAsByte, &dataAsInterface)
	if err != nil {
		return false, err
	}

	result := schema.Validate(dataAsInterface)
	if !result.IsValid() {
		details, err := json.MarshalIndent(result.ToList(), "", "  ")
		if err != nil {
			return false, err
		}
		fmt.Println(string(details))
		return false, nil
	}

	return true, nil
}

type VtmGameCharacter struct {
	Name        string       `json:"name"`
	Disciplines []Discipline `json:"disciplines"`
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
