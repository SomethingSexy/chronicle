package command

import (
	"context"
	"log"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewUpdateWorldCharacterCommand(persistence port.WorldPersistence) common.CommandHandler[corePort.UpdateWorldCharacter] {
	return updateWorldCharacterHandler{
		Persistence: persistence,
	}
}

type updateWorldCharacterHandler struct {
	Persistence port.WorldPersistence
}

// Patches an existing character that has already been added to the world
func (c updateWorldCharacterHandler) Handle(ctx context.Context, cmd corePort.UpdateWorldCharacter) error {
	log.Printf("Patching character %s to world %s", cmd.CharacterId, cmd.WorldId)

	return c.Persistence.UpsertCharacterToGameWorld(ctx, cmd.WorldId, cmd.CharacterId, &domain.WorldCharacter{
		Type: cmd.WorldCharacter.Type,
	})
}
