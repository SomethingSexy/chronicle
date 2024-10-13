package game

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

func NewGameRequest(g domain.Game) GameRequest {
	worlds := make([]*WorldRequest, len(g.Worlds))
	for x, world := range g.Worlds {
		worldRquest := NewWorldRequest(world)
		worlds[x] = &worldRquest
	}

	return GameRequest{
		ID:     g.GameId.String(),
		GameId: g.GameId.String(),
		Name:   g.Name,
		Type:   g.Type.String(),
		Worlds: worlds,
	}
}

type GameRequest struct {
	ID     string          `jsonapi:"primary,games"`
	GameId string          `jsonapi:"attr,gameId"`
	Name   string          `jsonapi:"attr,name"`
	Type   string          `jsonapi:"attr,type"`
	Worlds []*WorldRequest `jsonapi:"relation,worlds"`
}

func (a *GameRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required game fields")
	}

	if a.Type == "" {
		return errors.New("missing required game fields")
	}

	if _, err := uuid.Parse(a.GameId); err != nil {
		return errors.New("gameId must be valid UUID")
	}

	return nil
}

func (a *GameRequest) ToDomain() domain.Game {
	return domain.Game{
		GameId: uuid.MustParse(a.GameId),
		Name:   a.Name,
		Type:   domain.NewGameType(a.Type),
	}
}
