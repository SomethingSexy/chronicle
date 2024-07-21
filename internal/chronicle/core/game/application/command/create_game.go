package command

import (
	"context"
	"log"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
)

// TODO: Move this to a common spot
// TODO: Name this Command or CommandHandler
type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

type CreateGame struct {
	Game domain.Game
}

type CreateGameHander struct{}

func NewCreateGameCommand() CommandHandler[CreateGame] {

	return CreateGameHander{}
}

func (c CreateGameHander) Handle(ctx context.Context, cmd CreateGame) error {
	log.Printf("Create game %s", cmd.Game.Name)
	return nil
}
