package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication(persistence port.Persistence) port.ChronicleApplication {
	gameCommands := port.GameCommands{
		CreateGame:     command.NewCreateGameCommand(persistence.Game),
		CreateWorld:    command.NewCreateWorldCommand(persistence.Game),
		CreateLocation: command.NewCreateLocationCommand(persistence.Game),
	}

	characterCommands := port.CharacterCommands{
		CreateCharacter: command.NewCreateCharacterCommand(persistence.Character),
	}

	gameQueries := port.GameQueries{
		ListGames:     query.NewListGamesHandler(persistence.Game),
		GetGame:       query.NewGetGameHandler(persistence.Game),
		ListLocations: query.NewListLocationsHandler(persistence.Game),
		GetWorld:      query.NewGetWorldHandler(persistence.Game),
	}

	characterQueries := port.CharacterQueries{}

	return port.ChronicleApplication{
		Commands:    port.ChronicleCommands{GameCommands: gameCommands, CharacterCommands: characterCommands},
		Queries:     port.ChronicleQueries{GameQueries: gameQueries, CharacterQueries: characterQueries},
		Persistence: persistence,
	}
}
