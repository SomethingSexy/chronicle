package command

import (
	"context"
	"log"

	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewAddWorldCharacterCommand(persistence port.GamePersistence) common.CommandHandler[corePort.AddWorldCharacter] {
	return addWorldCharacterHandler{
		Persistence: persistence,
	}
}

type addWorldCharacterHandler struct {
	Persistence port.GamePersistence
}

// Adds a character to a world.  Further handling is created to update information about that link
func (c addWorldCharacterHandler) Handle(ctx context.Context, cmd corePort.AddWorldCharacter) error {
	log.Printf("Adding character %s to world %s", cmd.CharacterId, cmd.WorldId)

	return c.Persistence.AddCharacterToGameWorld(ctx, cmd.WorldId, cmd.CharacterId)
}
