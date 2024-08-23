package command

import (
	"context"
	"log"

	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewCreateWorldCommand(persistence port.ChronicleQueries) common.CommandHandler[gamePort.CreateWorld] {
	return createWorldHandler{
		Persistence: persistence,
	}
}

type createWorldHandler struct {
	Persistence port.ChronicleQueries
}

func (c createWorldHandler) Handle(ctx context.Context, cmd gamePort.CreateWorld) error {
	log.Printf("Creating world %s with id %s", cmd.World.Name, cmd.World.WorldId)
	_, err := c.Persistence.CreateWorld(ctx, cmd.World)
	if err != nil {
		return err
	}
	return nil
}
