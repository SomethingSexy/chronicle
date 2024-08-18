package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

func NewApplication() port.ChronicleApplication {
	commands := port.ChronicleCommands{
		CreateGame: command.NewCreateGameCommand(),
	}

	return port.ChronicleApplication{
		Commands: commands,
	}
}
