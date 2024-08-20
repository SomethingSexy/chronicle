package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
)

type CreateGame struct {
	Game domain.Game
}

type CreateGameHander common.CommandHandler[CreateGame]

type AllGamesQuery struct {
}
type ListGamesHandler common.QueryHandler[AllGamesQuery, []domain.Game]
