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
	WorldId uuid.UUID
}
type GetWorldHandler common.QueryHandler[GetWorldQuery, domain.World]

type AddWorldCharacter struct {
	WorldId     uuid.UUID
	CharacterId uuid.UUID
}
type AddWorldCharacterHandler common.CommandHandler[AddWorldCharacter]

type UpdateWorldCharacter struct {
	WorldId        uuid.UUID
	CharacterId    uuid.UUID
	WorldCharacter domain.WorldCharacter
}
type UpdateWorldCharacterHandler common.CommandHandler[UpdateWorldCharacter]
