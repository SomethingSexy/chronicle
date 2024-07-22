package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application/command"
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/port"
)

type GameApplication struct {
	Commands GameCommands
	Queries  GameQueries
	Server   corePort.HttpServer
}

type GameCommands struct {
	CreateGame command.CreateGameHander
}

type GameQueries struct {
}
