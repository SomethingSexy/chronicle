package game

import (
	"errors"
	"io"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/world"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

// Given a domain, return a GameRequest that will be used
// as a response payload.
func NewGameResponse(g domain.Game) GameRequest {
	worldRquest := world.NewWorldResponse(g.World)

	return GameRequest{
		ID:      g.GameId.String(),
		GameId:  g.GameId.String(),
		WorldId: g.WorldId.String(),
		Name:    g.Name,
		Type:    g.Type.String(),
		World:   &worldRquest,
	}
}

func NewGameRequest(body io.ReadCloser) (GameRequest, error) {
	var model GameRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

type GameRequest struct {
	ID      string              `jsonapi:"primary,games"`
	GameId  string              `jsonapi:"attr,gameId"`
	WorldId string              `jsonapi:"attr,worldId"`
	Name    string              `jsonapi:"attr,name"`
	Type    string              `jsonapi:"attr,type"`
	World   *world.WorldRequest `jsonapi:"relation,worlds"`
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

	if _, err := uuid.Parse(a.WorldId); err != nil {
		return errors.New("worldId must be valid UUID")
	}

	return nil
}

func (a *GameRequest) ToDomain() domain.Game {
	return domain.Game{
		GameId:  uuid.MustParse(a.GameId),
		WorldId: uuid.MustParse(a.WorldId),
		Name:    a.Name,
		Type:    domain.NewGameType(a.Type),
	}
}
