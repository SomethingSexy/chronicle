package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	worldPort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewGetWorldHandler(persistence port.WorldPersistence) worldPort.GetWorldHandler {
	return getWorldHandler{
		Persistence: persistence,
	}
}

type getWorldHandler struct {
	Persistence port.WorldPersistence
}

func (h getWorldHandler) Handle(ctx context.Context, q worldPort.GetWorldQuery) (domain.World, error) {
	world, err := h.Persistence.GetWorld(ctx, q.WorldId)
	if err != nil {
		return domain.World{}, err
	}

	// fetch locations tied to this world, later we can included/exclude these via jsonapi
	locations, err := h.Persistence.ListLocations(ctx, q.WorldId)
	if err != nil {
		return domain.World{}, err
	}
	world.Locations = locations

	characters, err := h.Persistence.ListCharacters(ctx, q.WorldId)
	if err != nil {
		return domain.World{}, err
	}
	world.Characters = characters

	return world, nil
}
