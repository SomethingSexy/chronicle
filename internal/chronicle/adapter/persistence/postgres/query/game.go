package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewGameQuery(queries *repository.Queries) GameQuery {
	return GameQuery{
		Queries: queries,
	}

}

type GameQuery struct {
	Queries *repository.Queries
}

func (g GameQuery) CreateGame(ctx context.Context, game domain.Game) (domain.Game, error) {
	// s := fmt.Sprintf("%x-%x-%x-%x-%x", myUUID.Bytes[0:4], myUUID.Bytes[4:6], myUUID.Bytes[6:8], myUUID.Bytes[8:10], myUUID.Bytes[10:16])

	args := repository.CreateGameParams{
		GameID: pgtype.UUID{
			Bytes: game.GameId,
			Valid: true,
		},
		Name: game.Name,
		Type: game.Type,
	}

	response, err := g.Queries.CreateGame(ctx, args)
	if err != nil {
		return domain.Game{}, err
	}

	return domain.Game{
		Name: response.Name,
		Type: response.Type,
	}, nil
}
