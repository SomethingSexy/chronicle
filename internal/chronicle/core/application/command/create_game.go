package command

import (
	"context"
	"log"

	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewCreateGameCommand(persistence port.ChronicleQueries) common.CommandHandler[gamePort.CreateGame] {

	return createGameHandler{
		Persistence: persistence,
	}
}

type createGameHandler struct {
	Persistence port.ChronicleQueries
}

func (c createGameHandler) Handle(ctx context.Context, cmd gamePort.CreateGame) error {
	log.Printf("Create game %s", cmd.Game.Name)
	_, err := c.Persistence.CreateGame(ctx, cmd.Game)
	if err != nil {
		return err
	}
	return nil
}
