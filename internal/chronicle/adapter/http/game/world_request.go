package game

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/character"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

func NewWorldRequest(w domain.World) WorldRequest {
	locations := make([]*LocationRequest, len(w.Locations))
	for i, location := range w.Locations {
		locationRequest := NewLocationRequest(location)
		locations[i] = &locationRequest
	}

	characters := make([]*character.CharacterRequest, len(w.Characters))
	for i, c := range w.Characters {
		characterRequest := character.NewCharacterRequest(c)
		characters[i] = &characterRequest
	}

	return WorldRequest{
		ID:         w.WorldId.String(),
		WorldId:    w.WorldId.String(),
		GameId:     w.GameId.String(),
		Name:       w.Name,
		Locations:  locations,
		Characters: characters,
	}
}

type WorldRequest struct {
	ID         string                        `jsonapi:"primary,worlds"`
	WorldId    string                        `jsonapi:"attr,worldId"`
	GameId     string                        `jsonapi:"attr,gameId"`
	Name       string                        `jsonapi:"attr,name"`
	Locations  []*LocationRequest            `jsonapi:"relation,locations"`
	Characters []*character.CharacterRequest `jsonapi:"relation,characters"`
}

func (a *WorldRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required game fields")
	}

	if _, err := uuid.Parse(a.GameId); err != nil {
		return errors.New("gameId must be valid UUID")
	}

	if _, err := uuid.Parse(a.WorldId); err != nil {
		return errors.New("worldId must be valid UUID")
	}

	return nil
}

func (a *WorldRequest) ToDomain() domain.World {
	return domain.World{
		WorldId: uuid.MustParse(a.WorldId),
		GameId:  uuid.MustParse(a.GameId),
		Name:    a.Name,
	}
}
