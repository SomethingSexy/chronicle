package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/adapter/http"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/port"
)

func NewApplication() port.GameApplication {
	commands := port.GameCommands{
		CreateGame: command.NewCreateGameCommand(),
	}

	return port.GameApplication{
		Commands: commands,
		Server:   http.NewHttpServer(commands, port.GameQueries{}),
	}
}
