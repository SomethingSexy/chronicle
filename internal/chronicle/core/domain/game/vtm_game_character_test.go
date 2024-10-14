package game_test

import (
	"testing"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain/game"
)

func TestGameCharacter_Validate_Valid(t *testing.T) {
	gameCharacter := domain.GameCharacter[game.VtmGameCharacter]{
		Data: game.VtmGameCharacter{
			Name: "John Doe",
		},
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
	gameCharacter := domain.GameCharacter[game.VtmGameCharacter]{
		Data: game.VtmGameCharacter{
			Name: "John Doe",
			Disciplines: []game.Discipline{{
				Name:  "Protean",
				Level: 1,
				Powers: []game.Power{{
					Name:        "Eyes of the Beast",
					Level:       1,
					Description: "See perfectly in total darkness with glowing red eyes.",
				}},
			}},
		},
	}

	valid, err := gameCharacter.Validate()

	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Error("should be valid")
	}
}
