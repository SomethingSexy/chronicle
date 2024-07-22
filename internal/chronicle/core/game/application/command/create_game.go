package command

import (
	"context"
	"log"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

type CreateGame struct {
	Game domain.Game
}

type CreateGameHander port.CommandHandler[CreateGame]

type createGameHandler struct{}

func NewCreateGameCommand() port.CommandHandler[CreateGame] {

	return createGameHandler{}
}

func (c createGameHandler) Handle(ctx context.Context, cmd CreateGame) error {
	log.Printf("Create game %s", cmd.Game.Name)
	return nil
}
