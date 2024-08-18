package command

import (
	"context"
	"log"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
)

type CreateGame struct {
	Game domain.Game
}

type CreateGameHander common.CommandHandler[CreateGame]

type createGameHandler struct{}

func NewCreateGameCommand() common.CommandHandler[CreateGame] {

	return createGameHandler{}
}

func (c createGameHandler) Handle(ctx context.Context, cmd CreateGame) error {
	log.Printf("Create game %s", cmd.Game.Name)
	return nil
}
