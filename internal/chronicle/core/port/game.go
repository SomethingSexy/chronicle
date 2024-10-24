package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/google/uuid"
)

type CreateGame struct {
	Game domain.Game
}

type CreateGameHander common.CommandHandler[CreateGame]

type AllGamesQuery struct {
}
type ListGamesHandler common.QueryHandler[AllGamesQuery, []domain.Game]

type GetGameQuery struct {
	GameId uuid.UUID
}
type GetGameHandler common.QueryHandler[GetGameQuery, domain.Game]

type UpdateGameCharacter struct {
	GameId      uuid.UUID
	CharacterId uuid.UUID
	Type        domain.CharacterType
	Character   map[string]interface{}
}
type UpdateGameCharacterHandler common.CommandHandler[UpdateGameCharacter]
