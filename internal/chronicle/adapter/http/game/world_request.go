package game

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

type WorldRequest struct {
	ID        string             `jsonapi:"primary,worlds"`
	WorldId   string             `jsonapi:"attr,worldId"`
	GameId    string             `jsonapi:"attr,gameId"`
	Name      string             `jsonapi:"attr,name"`
	Locations []*LocationRequest `jsonapi:"relation,locations"`
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
