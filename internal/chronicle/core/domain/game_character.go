package domain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kaptinlin/jsonschema"
)

// This should probably be generic, need to
// test this against the schema compiler though.
// Maybe after it is valid, we can marshall to a strict type
type GameCharacter struct {
	Data interface{}
}

func (g GameCharacter) Validate() (bool, error) {
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

	result := schema.Validate(g.Data)
	if !result.IsValid() {
		details, _ := json.MarshalIndent(result.ToList(), "", "  ")
		fmt.Println(string(details))
		return false, nil
	}

	return true, nil
}

type VtmGameCharacter struct {
	Name string `json:"name"`
}
