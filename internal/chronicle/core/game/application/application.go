package application

import "github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application/command"

type GameApplication struct {
	Commands GameCommands
	Queries  GameQueries
}

type GameCommands struct {
	CreateGame command.CreateGameHander
}

type GameQueries struct {
}

// TODO: This should have a New function to setup the commands, repositories
// and routes for this domain
