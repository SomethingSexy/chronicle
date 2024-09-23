package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	worldPort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewGetWorldHandler(persistence port.ChronicleQueries) worldPort.GetWorldHandler {
	return getWorldHandler{
		Persistence: persistence,
	}
}

type getWorldHandler struct {
	Persistence port.ChronicleQueries
}

func (h getWorldHandler) Handle(ctx context.Context, q worldPort.GetWorldQuery) (domain.World, error) {
	world, err := h.Persistence.GetWorld(ctx, q.GameId, q.WorldId)
	if err != nil {
		return domain.World{}, err
	}

	// fetch locations tied to this world, later we can included/exclude these via jsonapi
	locations, err := h.Persistence.ListLocations(ctx, q.GameId, q.WorldId)
	if err != nil {
		return domain.World{}, err
	}

	world.Loocation = locations

	return world, nil
}
