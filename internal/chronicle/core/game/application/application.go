package application

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/adapter"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/port"
)

func NewApplication() port.GameApplication {
	commands := port.GameCommands{
		CreateGame: command.NewCreateGameCommand(),
	}

	return port.GameApplication{
		Commands: commands,
		Server:   adapter.NewHttpServer(commands, port.GameQueries{}),
	}
}

// func (a GameApplication) Server() port.HttpServer {
// 	return adapter.NewHttpServer(a)
// }
