package query

import (
	"context"
	"log"
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
	world, err := g.Queries.GetWorldFromUuid(ctx, game.WorldId)
	if err != nil {
		return domain.Game{}, err
	}

	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	args := repository.CreateGameParams{
		GameID:    game.GameId,
		WorldID:   world.ID,
		Name:      game.Name,
		Type:      game.Type.String(),
		CreatedAt: ts,
		UpdatedAt: ts,
	}

	response, err := g.Queries.CreateGame(ctx, args)
	if err != nil {
		return domain.Game{}, err
	}

	return domain.Game{
		Name: response.Name,
		Type: domain.NewGameType(response.Type),
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
			Type:   domain.NewGameType(game.Type),
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
	world, err := g.Queries.GetWorld(ctx, response.WorldID)
	if err != nil {
		return domain.Game{}, err
	}

	log.Println(response.Type, domain.NewGameType(response.Type))

	game := domain.Game{
		Name:    response.Name,
		Type:    domain.NewGameType(response.Type),
		GameId:  response.GameID,
		WorldId: world.WorldID,
	}

	return game, nil
}

func (g GameQuery) GetGameWorld(ctx context.Context, gameId uuid.UUID) (domain.World, error) {
	response, err := g.Queries.GetGameWorld(ctx, gameId)
	if err != nil {
		return domain.World{}, err
	}

	world := domain.World{
		WorldId: response.WorldID,
		Name:    response.Name,
	}

	return world, nil
}
