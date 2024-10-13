package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/google/uuid"
)

type CreateLocation struct {
	Location domain.Location
}
type CreateLocationHander common.CommandHandler[CreateLocation]

type LocationsQuery struct {
	WorldId uuid.UUID
}
type ListLocationsHandler common.QueryHandler[LocationsQuery, []domain.Location]
