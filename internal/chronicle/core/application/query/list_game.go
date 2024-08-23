package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewListGamesHandler(persistence port.ChronicleQueries) gamePort.ListGamesHandler {
	return listGamesHandler{
		Persistence: persistence,
	}
}

type listGamesHandler struct {
	Persistence port.ChronicleQueries
}

func (h listGamesHandler) Handle(ctx context.Context, _ gamePort.AllGamesQuery) ([]domain.Game, error) {
	// Fetching these separately for now, when necessary we could make these joins at the lower level but
	// this gets me more power to decide what I want to fetch when
	games, err := h.Persistence.ListGames(ctx)
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		world, err := h.Persistence.GetGameWorlds(ctx, game.GameId)
		if err != nil {
			return nil, err
		}
		// TODO: Need to decide how many worlds to support in a game
		if len(world) > 0 {
			game.World = &world[0]
		}
	}

	return games, nil
}
