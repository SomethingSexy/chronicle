package command

import (
	"context"
	"log"

	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
)

func NewCreateCharacterCommand(persistence port.CharacterPersistence) common.CommandHandler[corePort.CreateCharacter] {
	return createCharacterHandler{
		Persistence: persistence,
	}
}

type createCharacterHandler struct {
	Persistence port.CharacterPersistence
}

func (c createCharacterHandler) Handle(ctx context.Context, cmd corePort.CreateCharacter) error {
	log.Printf("Creating character %s with id %s", cmd.Character.Name, cmd.Character.CharacterId)
	_, err := c.Persistence.CreateCharacter(ctx, cmd.Character)
	if err != nil {
		return err
	}
	return nil
}
