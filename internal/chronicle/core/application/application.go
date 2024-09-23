package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication(persistence port.ChronicleQueries) port.ChronicleApplication {
	commands := port.ChronicleCommands{
		CreateGame:     command.NewCreateGameCommand(persistence),
		CreateWorld:    command.NewCreateWorldCommand(persistence),
		CreateLocation: command.NewCreateLocationCommand(persistence),
	}

	queries := port.GameQueries{
		ListGames:     query.NewListGamesHandler(persistence),
		GetGame:       query.NewGetGameHandler(persistence),
		ListLocations: query.NewListLocationsHandler(persistence),
		GetWorld:      query.NewGetWorldHandler(persistence),
	}

	return port.ChronicleApplication{
		Commands:    commands,
		Queries:     queries,
		Persistence: persistence,
	}
}
