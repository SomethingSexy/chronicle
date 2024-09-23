package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/google/uuid"
)

type CreateWorld struct {
	World domain.World
}
type CreateWorldHander common.CommandHandler[CreateWorld]

type GetWorldQuery struct {
	GameId  uuid.UUID
	WorldId uuid.UUID
}
type GetWorldHandler common.QueryHandler[GetWorldQuery, domain.World]
