package port

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
)

type GameService interface {
	CreateUser(ctx context.Context, game domain.Game) error
}
