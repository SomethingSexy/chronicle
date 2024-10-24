package domain_test

import (
	"testing"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
)

func TestGameCharacter_Validate_Valid(t *testing.T) {
	// gameCharacter := domain.GameCharacter[domain.VtmGameCharacter]{
	// 	Data: domain.VtmGameCharacter{
	// 		Name: "John Doe",
	// 	},
	// }
	gameCharacter := domain.NewVtmGameCharacter(domain.VTM, map[string]interface{}{
		"name": "John Doe",
	})

	valid, err := domain.Validate(gameCharacter)

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("should be valid")
	}
}

func TestGameCharacter_Validate_Disciplines_Valid(t *testing.T) {
	// gameCharacter := domain.GameCharacter[domain.VtmGameCharacter]{
	// 	Data: domain.VtmGameCharacter{
	// 		Name: "John Doe",
	// 		Disciplines: []domain.Discipline{{
	// 			Name:  "Protean",
	// 			Level: 1,
	// 			Powers: []domain.Power{{
	// 				Name:        "Eyes of the Beast",
	// 				Level:       1,
	// 				Description: "See perfectly in total darkness with glowing red eyes.",
	// 			}},
	// 		}},
	// 	},
	// }

	gameCharacter := domain.NewVtmGameCharacter(domain.VTM, map[string]interface{}{
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
		t.Error("should be valid")
	}
}
