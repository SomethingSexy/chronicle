package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
)

func TestGameCharacter_Validate_Valid(t *testing.T) {
	var data interface{}

	if err := json.Unmarshal([]byte(`{
		"name": "John Doe"
	}`), &data); err != nil {
		t.Fatalf("Failed to unmarshal test cases: %v", err)
	}

	gameCharacter := domain.GameCharacter{
		Data: data,
	}

	valid, err := gameCharacter.Validate()

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("should be valid")
	}
}

func TestGameCharacter_Validate_Disciplines_Valid(t *testing.T) {
	// {
	// 	"name": "Protean",
	// 	"level": 3,
	// 	"powers": [
	// 		{
	// 			"name": "Eyes of the Beast",
	// 			"level": 1,
	// 			"description": "See perfectly in total darkness with glowing red eyes."
	// 		},
	// 		{
	// 			"name": "Feral Weapons",
	// 			"level": 2,
	// 			"description": "Grow claws or fangs, gaining lethal unarmed attacks."
	// 		},
	// 		{
	// 			"name": "Metamorphosis",
	// 			"level": 3,
	// 			"description": "Transform your body to gain animal-like characteristics."
	// 		}
	// 	]
	// },
	var data interface{}

	if err := json.Unmarshal([]byte(`{
		"name": "John Doe",
		"disciplines": [{
		  "name": "Protean",
			"level": 1,
			"powers": [{
	 			"name": "Eyes of the Beast",
	 			"description": "See perfectly in total darkness with glowing red eyes."
	 		}]
		}]
	}`), &data); err != nil {
		t.Fatalf("Failed to unmarshal test cases: %v", err)
	}

	gameCharacter := domain.GameCharacter{
		Data: data,
	}

	valid, err := gameCharacter.Validate()

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("should be valid")
	}
}
