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
	return h.Persistence.ListGames(ctx)
}
