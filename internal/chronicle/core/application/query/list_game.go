package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewListGamesHandler(persistence port.Persistence) gamePort.ListGamesHandler {
	return listGamesHandler{
		Persistence: persistence,
	}
}

type listGamesHandler struct {
	Persistence port.Persistence
}

func (h listGamesHandler) Handle(ctx context.Context, _ gamePort.AllGamesQuery) ([]domain.Game, error) {
	// Fetching these separately for now, when necessary we could make these joins at the lower level
	// For now, this gives me more power to move this logic around as needed
	games, err := h.Persistence.Game.ListGames(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(games); i++ {
		world, err := h.Persistence.Game.GetGameWorld(ctx, games[i].GameId)
		if err != nil {
			return nil, err
		}

		characters, err := h.Persistence.World.ListCharacters(ctx, world.WorldId)
		if err != nil {
			return nil, err
		}
		world.Characters = characters

		games[i].World = world
	}

	return games, nil
}
