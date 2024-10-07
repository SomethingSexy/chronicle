package port

import (
	"context"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

// Port for the rest of the application to interface
// with the persistence layer
type Persistence struct {
	Game      GamePersistence
	Character CharacterPersistence
}

type GamePersistence interface {
	CreateGame(ctx context.Context, game domain.Game) (domain.Game, error)
	ListGames(ctx context.Context) ([]domain.Game, error)
	GetGame(ctx context.Context, id uuid.UUID) (domain.Game, error)
	// TODO: This should turn into a ListWorlds with a filter by gameId
	GetGameWorlds(ctx context.Context, gameId uuid.UUID) ([]domain.World, error)
	CreateWorld(ctx context.Context, world domain.World) (domain.World, error)
	GetWorld(ctx context.Context, gameId uuid.UUID, worldId uuid.UUID) (domain.World, error)

	CreateLocation(ct context.Context, location domain.Location) (domain.Location, error)
	ListLocations(ctx context.Context, gameId uuid.UUID, worldId uuid.UUID) ([]domain.Location, error)

	AddCharacterToGameWorld(ctx context.Context, worldId uuid.UUID, characterId uuid.UUID) error
}

type CharacterPersistence interface {
	CreateCharacter(ctx context.Context, character domain.Character) (domain.Character, error)
}
