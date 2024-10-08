package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewGetGameHandler(persistence port.GamePersistence) gamePort.GetGameHandler {
	return getGameHandler{
		Persistence: persistence,
	}
}

type getGameHandler struct {
	Persistence port.GamePersistence
}

func (h getGameHandler) Handle(ctx context.Context, q gamePort.GetGameQuery) (domain.Game, error) {
	return h.Persistence.GetGame(ctx, q.GameId)
}
