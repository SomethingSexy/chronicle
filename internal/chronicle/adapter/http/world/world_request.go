package world

import (
	"errors"
	"io"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/character"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewWorldResponse(w domain.World) WorldRequest {
	locations := make([]*LocationRequest, len(w.Locations))
	for i, location := range w.Locations {
		locationRequest := NewLocationResponse(location)
		locations[i] = &locationRequest
	}

	characters := make([]*character.CharacterRequest, len(w.Characters))
	for i, c := range w.Characters {
		characterRequest := character.NewCharacterResponse(c)
		characters[i] = &characterRequest
	}

	return WorldRequest{
		ID:         w.WorldId.String(),
		WorldId:    w.WorldId.String(),
		Name:       w.Name,
		Locations:  locations,
		Characters: characters,
	}
}

func NewWorldRequest(body io.ReadCloser) (WorldRequest, error) {
	var model WorldRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

type WorldRequest struct {
	ID         string                        `jsonapi:"primary,worlds"`
	WorldId    string                        `jsonapi:"attr,worldId"`
	Name       string                        `jsonapi:"attr,name"`
	Locations  []*LocationRequest            `jsonapi:"relation,locations"`
	Characters []*character.CharacterRequest `jsonapi:"relation,characters"`
}

func (a *WorldRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required game fields")
	}

	if _, err := uuid.Parse(a.WorldId); err != nil {
		return errors.New("worldId must be valid UUID")
	}

	return nil
}

func (a *WorldRequest) ToDomain() domain.World {
	return domain.World{
		WorldId: uuid.MustParse(a.WorldId),
		Name:    a.Name,
	}
}
