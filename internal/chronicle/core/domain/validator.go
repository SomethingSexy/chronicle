package domain

import (
	"log"

	"github.com/kaptinlin/jsonschema"
)

type Validator interface {
	Data() any
	Schema() ([]byte, error)
}

// General validation function that process a struct
// and validates it against a schema.
//
// TODO: This could just return an error
func Validate(t Validator) (bool, error) {
	typeSchema, err := t.Schema()
	if err != nil {
		return false, err
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(typeSchema)
	if err != nil {
		return false, err
	}

	// Need to figure this out but I believe the validator
	// requires it to be a map of interfaces.
	// For now we are going to marshall and unmarshall back into
	// a generic interface of maps so we can validate the schema
	// Hide it all here
	// dataAsByte, err := json.Marshal(t.Data)
	// if err != nil {
	// 	return false, err
	// }

	// var dataAsInterface interface{}
	// err = json.Unmarshal(dataAsByte, &dataAsInterface)
	// if err != nil {
	// 	return false, err
	// }

	result := schema.Validate(t.Data())
	if !result.IsValid() {
		log.Println(result)
		// TODO: Need to figure out what we are going to actually return here
		// details, err := json.MarshalIndent(result.ToList(), "", "  ")
		// if err != nil {
		// 	return false, err
		// }
		// fmt.Println(string(details))
		return false, nil
	}

	return true, nil
}
