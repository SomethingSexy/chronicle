package domain

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/kaptinlin/jsonschema"
)

// This should probably be generic, need to
// test this against the schema compiler though.
// Maybe after it is valid, we can marshall to a strict type
type GameCharacter[D Validator] struct {
	Data D
	Type GameType
}

func (g GameCharacter[D]) Validate() (bool, error) {
	characterSchema, err := g.Data.Schema()
	if err != nil {
		return false, err
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(characterSchema)
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
		log.Println(result)
		// details, err := json.MarshalIndent(result.ToList(), "", "  ")
		// if err != nil {
		// 	return false, err
		// }
		// fmt.Println(string(details))
		return false, nil
	}

	return true, nil
}
