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

	characters, err := h.Persistence.World.ListCharacters(ctx, world.WorldId)
	if err != nil {
		return domain.Game{}, err
	}
	world.Characters = characters

	game.World = world

	return game, nil
}
