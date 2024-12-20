package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewListLocationsHandler(persistence port.WorldPersistence) gamePort.ListLocationsHandler {
	return listLocationsHandler{
		Persistence: persistence,
	}
}

type listLocationsHandler struct {
	Persistence port.WorldPersistence
}

func (h listLocationsHandler) Handle(ctx context.Context, query gamePort.LocationsQuery) ([]domain.Location, error) {
	locations, err := h.Persistence.ListLocations(ctx, query.WorldId)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
