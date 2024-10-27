package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewGetGameHandler(persistence port.Persistence) gamePort.GetGameHandler {
	return getGameHandler{
		Persistence: persistence,
	}
}

type getGameHandler struct {
	Persistence port.Persistence
}

func (h getGameHandler) Handle(ctx context.Context, q gamePort.GetGameQuery) (domain.Game, error) {
	game, err := h.Persistence.Game.GetGame(ctx, q.GameId)
	if err != nil {
		return domain.Game{}, err
	}

	world, err := h.Persistence.Game.GetGameWorld(ctx, game.GameId)
	if err != nil {
		return domain.Game{}, err
	}

	// TODO: In the future, probably makes more sense to combine world character fetch and
	// game-character fetch into a single entity
	worldCharacters, err := h.Persistence.World.ListCharacters(ctx, world.WorldId)
	if err != nil {
		return domain.Game{}, err
	}
	world.Characters = worldCharacters
	game.World = world

	characters, err := h.Persistence.Game.ListCharacters(ctx, game.GameId)
	if err != nil {
		return domain.Game{}, err
	}
	game.Characters = characters

	return game, nil
}
