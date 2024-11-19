package domain_test

import (
	"testing"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

func TestGameCharacter_Validate_Valid(t *testing.T) {
	gameCharacter := domain.NewVtmGameCharacter(uuid.New(), uuid.New(), domain.VTM, map[string]interface{}{
		"name": "John Doe",
	})

	valid, err := domain.Validate(gameCharacter)

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("Character should be valid")
	}
}

func TestGameCharacter_Validate_Disciplines_Valid(t *testing.T) {
	gameCharacter := domain.NewVtmGameCharacter(uuid.New(), uuid.New(), domain.VTM, map[string]interface{}{
		"name": "John Doe",
		"dsiciplines": []map[string]interface{}{{
			"name":  "Protean",
			"level": 1,
			"powers": []map[string]interface{}{{
				"name":        "Eyes of the Beast",
				"level":       1,
				"Description": "See perfectly in total darkness with glowing red eyes.",
			}},
		}},
	})

	valid, err := domain.Validate(gameCharacter)

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("Character should be valid")
	}
}
