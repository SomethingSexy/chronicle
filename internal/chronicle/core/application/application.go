package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication(persistence port.Persistence) port.ChronicleApplication {
	gameCommands := port.GameCommands{
		CreateGame:          command.NewCreateGameCommand(persistence.Game),
		UpdateGameCharacter: command.NewUpdateGameCharacterCommand(persistence),
	}

	characterCommands := port.CharacterCommands{
		CreateCharacter: command.NewCreateCharacterCommand(persistence.Character),
	}

	gameQueries := port.GameQueries{
		ListGames: query.NewListGamesHandler(persistence),
		GetGame:   query.NewGetGameHandler(persistence),
	}

	characterQueries := port.CharacterQueries{}

	worldQueries := port.WorldQueries{
		ListLocations: query.NewListLocationsHandler(persistence.World),
		GetWorld:      query.NewGetWorldHandler(persistence.World),
	}

	worldCommands := port.WorldCommands{
		CreateWorld:       command.NewCreateWorldCommand(persistence.World),
		CreateLocation:    command.NewCreateLocationCommand(persistence.World),
		AddWorldCharacter: command.NewAddWorldCharacterCommand(persistence.World),
	}

	return port.ChronicleApplication{
		Commands: port.ChronicleCommands{
			GameCommands: gameCommands, CharacterCommands: characterCommands, WorldCommands: worldCommands,
		},
		Queries: port.ChronicleQueries{
			GameQueries: gameQueries, CharacterQueries: characterQueries, WorldQueries: worldQueries,
		},
		Persistence: persistence,
	}
}
