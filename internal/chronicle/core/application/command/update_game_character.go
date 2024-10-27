package command

import (
	"context"
	"errors"
	"log"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/google/uuid"
)

func NewUpdateGameCharacterCommand(persistence port.Persistence) common.CommandHandler[corePort.UpdateGameCharacter] {
	return updateGameCharacterHandler{
		Persistence: persistence,
	}
}

type updateGameCharacterHandler struct {
	Persistence port.Persistence
}

// Allows game specific additions to a character that is tied to a world.
func (c updateGameCharacterHandler) Handle(ctx context.Context, cmd corePort.UpdateGameCharacter) error {
	game, err := c.Persistence.Game.GetGame(ctx, cmd.GameId)
	if err != nil {
		return err
	}

	gameCharacter := domain.NewGameCharacter(
		game.Type,
		// Nil type for now because during the update we don't care about this, really only
		// want it on the response...tODO - maybe better way to handle this
		uuid.Nil,
		cmd.CharacterId,
		cmd.Type,
		cmd.Character,
	)

	valid, err := domain.Validate(gameCharacter)
	if err != nil {
		return err
	}

	if !valid {
		// TODO: Once we figure out how we are handling errors here
		// return that appropriately
		return errors.New("Character is not valid")
	}

	log.Printf("Updating character %s to game %s", cmd.CharacterId, cmd.GameId)

	return c.Persistence.Game.UpdateCharacter(ctx, cmd.GameId, cmd.CharacterId, gameCharacter)
}
