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

	worlds, err := h.Persistence.Game.GetGameWorlds(ctx, game.GameId)
	if err != nil {
		return domain.Game{}, err
	}

	for x := 0; x < len(worlds); x++ {
		characters, err := h.Persistence.World.ListCharacters(ctx, worlds[x].GameId, worlds[x].WorldId)
		if err != nil {
			return domain.Game{}, err
		}
		worlds[x].Characters = characters
	}

	game.Worlds = worlds

	return game, nil
}
