package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication(persistence port.ChronicleQueries) port.ChronicleApplication {
	commands := port.ChronicleCommands{
		CreateGame: command.NewCreateGameCommand(persistence),
	}

	return port.ChronicleApplication{
		Commands:    commands,
		Persistence: persistence,
	}
}
