package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication(persistence port.Persistence) port.ChronicleApplication {
	commands := port.ChronicleCommands{
		CreateGame:     command.NewCreateGameCommand(persistence.Game),
		CreateWorld:    command.NewCreateWorldCommand(persistence.Game),
		CreateLocation: command.NewCreateLocationCommand(persistence.Game),
	}

	queries := port.GameQueries{
		ListGames:     query.NewListGamesHandler(persistence.Game),
		GetGame:       query.NewGetGameHandler(persistence.Game),
		ListLocations: query.NewListLocationsHandler(persistence.Game),
		GetWorld:      query.NewGetWorldHandler(persistence.Game),
	}

	return port.ChronicleApplication{
		Commands:    commands,
		Queries:     queries,
		Persistence: persistence,
	}
}
