package port

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

// TODO: Probably a better name for this
type ChronicleQueries interface {
	CreateGame(ctx context.Context, game domain.Game) (domain.Game, error)
	ListGames(ctx context.Context) ([]domain.Game, error)
	GetGame(ctx context.Context, id uuid.UUID) (domain.Game, error)
}
