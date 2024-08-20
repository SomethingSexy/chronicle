package query

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
)

func NewGameQuery(queries *repository.Queries) GameQuery {
	return GameQuery{
		Queries: queries,
	}

}

type GameQuery struct {
	Queries *repository.Queries
}

// Create a game
func (g GameQuery) CreateGame(ctx context.Context, game domain.Game) (domain.Game, error) {
	args := repository.CreateGameParams{
		GameID: game.GameId,
		Name:   game.Name,
		Type:   game.Type,
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

// List all available games.
// TODO:
//   - Limit games by user who created them
//   - Limit games if they are marked private
//   - Allow admin to see all games
func (g GameQuery) ListGames(ctx context.Context) ([]domain.Game, error) {
	response, err := g.Queries.ListGames(ctx)
	if err != nil {
		return nil, err
	}

	games := make([]domain.Game, len(response))

	for i, game := range response {
		games[i] = domain.Game{
			Name:   game.Name,
			Type:   game.Type,
			GameId: game.GameID,
		}
	}

	return games, nil
}
