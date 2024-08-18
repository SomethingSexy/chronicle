package port

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
)

// TODO: Probably a better name for this
type ChronicleQueries interface {
	CreateGame(ctx context.Context, game domain.Game) (domain.Game, error)
}
