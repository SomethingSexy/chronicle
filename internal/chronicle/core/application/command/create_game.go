package command

import (
	"context"
	"log"

	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewCreateGameCommand(persistence port.GamePersistence) common.CommandHandler[gamePort.CreateGame] {
	return createGameHandler{
		Persistence: persistence,
	}
}

type createGameHandler struct {
	Persistence port.GamePersistence
}

func (c createGameHandler) Handle(ctx context.Context, cmd gamePort.CreateGame) error {
	log.Printf("Creating game %s with id %s", cmd.Game.Name, cmd.Game.GameId)
	_, err := c.Persistence.CreateGame(ctx, cmd.Game)
	if err != nil {
		return err
	}
	return nil
}
