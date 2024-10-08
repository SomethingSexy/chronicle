package query

import (
	"context"
	"time"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// TODO: Split this when it gets too big
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
	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	args := repository.CreateGameParams{
		GameID:    game.GameId,
		Name:      game.Name,
		Type:      game.Type,
		CreatedAt: ts,
		UpdatedAt: ts,
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

func (g GameQuery) GetGame(ctx context.Context, id uuid.UUID) (domain.Game, error) {
	response, err := g.Queries.GetGameFromUuid(ctx, id)
	if err != nil {
		return domain.Game{}, err
	}

	game := domain.Game{
		Name:   response.Name,
		Type:   response.Type,
		GameId: response.GameID,
	}

	return game, nil
}

func (g GameQuery) GetGameWorlds(ctx context.Context, gameId uuid.UUID) ([]domain.World, error) {
	response, err := g.Queries.GetGameWorlds(ctx, gameId)
	if err != nil {
		return nil, err
	}

	worlds := make([]domain.World, len(response))

	for i, world := range response {
		worlds[i] = domain.World{
			WorldId: world.WorldID,
			GameId:  gameId,
			Name:    world.Name,
		}
	}

	return worlds, nil
}
