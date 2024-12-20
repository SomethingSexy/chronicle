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
	World     WorldPersistence
}

type GamePersistence interface {
	CreateGame(ctx context.Context, game domain.Game) (domain.Game, error)
	ListGames(ctx context.Context) ([]domain.Game, error)
	GetGame(ctx context.Context, id uuid.UUID) (domain.Game, error)
	GetGameWorld(ctx context.Context, gameId uuid.UUID) (domain.World, error)
	UpdateCharacter(ctx context.Context, gameId uuid.UUID, characterId uuid.UUID, character domain.GameCharacter) error
	ListCharacters(ctx context.Context, gameId uuid.UUID) ([]domain.GameCharacter, error)
}

type CharacterPersistence interface {
	CreateCharacter(ctx context.Context, character domain.Character) (domain.Character, error)
}

type WorldPersistence interface {
	CreateWorld(ctx context.Context, world domain.World) (domain.World, error)
	GetWorld(ctx context.Context, worldId uuid.UUID) (domain.World, error)

	CreateLocation(ct context.Context, location domain.Location) (domain.Location, error)
	ListLocations(ctx context.Context, worldId uuid.UUID) ([]domain.Location, error)

	ListCharacters(ctx context.Context, worldId uuid.UUID) ([]domain.Character, error)
	AddCharacterToGameWorld(ctx context.Context, worldId uuid.UUID, characterId uuid.UUID) error
}
