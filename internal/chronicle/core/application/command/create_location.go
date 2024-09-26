package command

import (
	"context"
	"log"

	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewCreateLocationCommand(persistence port.GamePersistence) common.CommandHandler[gamePort.CreateLocation] {
	return createLocationHandler{
		Persistence: persistence,
	}
}

type createLocationHandler struct {
	Persistence port.GamePersistence
}

func (c createLocationHandler) Handle(ctx context.Context, cmd gamePort.CreateLocation) error {
	log.Printf("Creating location %s with id %s", cmd.Location.Name, cmd.Location.LocationId)
	_, err := c.Persistence.CreateLocation(ctx, cmd.Location)
	if err != nil {
		return err
	}
	return nil
}
