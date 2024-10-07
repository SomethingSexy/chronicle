package query

import (
	"context"
	"strings"
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

func (g GameQuery) CreateWorld(ctx context.Context, world domain.World) (domain.World, error) {
	game, err := g.Queries.GetGameFromUuid(ctx, world.GameId)
	if err != nil {
		return domain.World{}, err
	}

	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	args := repository.CreateWorldParams{
		WorldID:   world.WorldId,
		Name:      world.Name,
		GameID:    game.ID,
		CreatedAt: ts,
		UpdatedAt: ts,
	}

	response, err := g.Queries.CreateWorld(ctx, args)
	if err != nil {
		return domain.World{}, err
	}

	return domain.World{
		WorldId: response.WorldID,
		Name:    response.Name,
	}, nil
}

func (g GameQuery) GetWorld(ctx context.Context, gameId uuid.UUID, worldId uuid.UUID) (domain.World, error) {
	world, err := g.Queries.GetWorldFromUuid(ctx, worldId)
	if err != nil {
		return domain.World{}, err
	}

	return domain.World{
		WorldId: world.WorldID,
		GameId:  gameId,
		Name:    world.Name,
	}, nil
}

func (g GameQuery) CreateLocation(ctx context.Context, location domain.Location) (domain.Location, error) {
	world, err := g.Queries.GetWorldFromUuid(ctx, location.WorldId)
	if err != nil {
		return domain.Location{}, err
	}

	path := ""
	pathsLength := len(location.Path)

	if pathsLength > 0 {
		if pathsLength == 1 {
			path = location.Path[0].String()

		} else {
			for _, part := range location.Path {
				path = path + "." + part.String()
			}
		}
	}

	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	args := repository.CreateLocationParams{
		LocationID: location.LocationId,
		WorldID:    world.ID,
		GameID:     world.GameID,
		Name:       location.Name,
		Type:       location.Type,
		Path: pgtype.Text{
			String: path,
			Valid:  true,
		},
		CreatedAt: ts,
		UpdatedAt: ts,
	}

	_, err = g.Queries.CreateLocation(ctx, args)
	if err != nil {
		return domain.Location{}, err
	}

	return location, nil
}

func (g GameQuery) ListLocations(ctx context.Context, gameId uuid.UUID, worldId uuid.UUID) ([]domain.Location, error) {
	response, err := g.Queries.GetWorldLocations(ctx, repository.GetWorldLocationsParams{
		GameID:  gameId,
		WorldID: worldId,
	})
	if err != nil {
		return []domain.Location{}, err
	}

	locations := make([]domain.Location, len(response))

	for i, location := range response {
		var paths []uuid.UUID
		// Split might return a string with a single element here
		if location.Path.String != "" {
			rawPaths := strings.Split(location.Path.String, ".")
			for _, path := range rawPaths {
				// This error should never but ya know...
				parsed, err := uuid.Parse(path)
				if err != nil {
					return []domain.Location{}, err
				}
				paths = append(paths, parsed)
			}
		}

		locations[i] = domain.Location{
			LocationId: location.LocationID,
			WorldId:    location.WorldID_2,
			Name:       location.Name,
			Type:       location.Type,
			Path:       paths,
		}
	}

	return locations, nil
}

func (g GameQuery) AddCharacterToGameWorld(ctx context.Context, worldId uuid.UUID, characterId uuid.UUID) error {
	world, err := g.Queries.GetWorldFromUuid(ctx, worldId)
	if err != nil {
		return err
	}

	character, err := g.Queries.GetCharacterFromUuid(ctx, characterId)
	if err != nil {
		return err
	}

	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}

	return g.Queries.AddCharacterToGameWorld(ctx, repository.AddCharacterToGameWorldParams{
		// Just generating this now for future use but since we are using a unique
		// index for worldId and characterId, this is probably too important
		WorldCharacterID: uuid.New(),
		WorldID:          world.ID,
		CharacterID:      character.ID,
		UpdatedAt:        ts,
		CreatedAt:        ts,
	})
}

func (g GameQuery) ListCharacters(ctx context.Context, gameId uuid.UUID, worldId uuid.UUID) ([]domain.Character, error) {
	response, err := g.Queries.GetWorldCharacters(ctx, worldId)
	if err != nil {
		return nil, err
	}

	characters := make([]domain.Character, len(response))

	for i, character := range response {
		characters[i] = domain.Character{
			CharacterId: character.CharacterID,
			Name:        character.Name,
			Description: character.Description.String,
		}
	}

	return characters, nil
}
