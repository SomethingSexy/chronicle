package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application/command"
)

type ChronicleApplication struct {
	Commands ChronicleCommands
	Queries  GameQueries
}

type ChronicleCommands struct {
	CreateGame command.CreateGameHander
}

type GameQueries struct {
}
