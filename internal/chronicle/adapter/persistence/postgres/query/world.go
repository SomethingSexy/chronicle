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
func NewWorldQuery(queries *repository.Queries) WorldQuery {
	return WorldQuery{
		Queries: queries,
	}
}

type WorldQuery struct {
	Queries *repository.Queries
}

func (g WorldQuery) CreateWorld(ctx context.Context, world domain.World) (domain.World, error) {
	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	args := repository.CreateWorldParams{
		WorldID:   world.WorldId,
		Name:      world.Name,
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

func (g WorldQuery) GetWorld(ctx context.Context, worldId uuid.UUID) (domain.World, error) {
	world, err := g.Queries.GetWorldFromUuid(ctx, worldId)
	if err != nil {
		return domain.World{}, err
	}

	return domain.World{
		WorldId: world.WorldID,
		Name:    world.Name,
	}, nil
}

func (g WorldQuery) CreateLocation(ctx context.Context, location domain.Location) (domain.Location, error) {
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

func (g WorldQuery) ListLocations(ctx context.Context, worldId uuid.UUID) ([]domain.Location, error) {
	response, err := g.Queries.GetWorldLocations(ctx, worldId)
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

func (g WorldQuery) UpsertCharacterToGameWorld(ctx context.Context, worldId uuid.UUID, characterId uuid.UUID, character *domain.WorldCharacter) error {
	world, err := g.Queries.GetWorldFromUuid(ctx, worldId)
	if err != nil {
		return err
	}

	existingCharacter, err := g.Queries.GetCharacterFromUuid(ctx, characterId)
	if err != nil {
		return err
	}

	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}

	requestArgs := repository.UpsertCharacterToGameWorldParams{
		// Just generating this now for future use but since we are using a unique
		// index for worldId and characterId, this is probably too important
		WorldCharacterID: uuid.New(),
		WorldID:          world.ID,
		CharacterID:      existingCharacter.ID,
		UpdatedAt:        ts,
		CreatedAt:        ts,
	}

	if character != nil {
		requestArgs.CharacterType = pgtype.Text{
			Valid:  true,
			String: character.Type.String(),
		}
	}

	return g.Queries.UpsertCharacterToGameWorld(ctx, requestArgs)
}

func (g WorldQuery) ListCharacters(ctx context.Context, worldId uuid.UUID) ([]domain.Character, error) {
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
