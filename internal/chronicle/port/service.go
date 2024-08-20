package port

import (
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
)

type ChronicleApplication struct {
	Commands    ChronicleCommands
	Queries     GameQueries
	Persistence ChronicleQueries
}

type ChronicleCommands struct {
	CreateGame corePort.CreateGameHander
}

type GameQueries struct {
	ListGames corePort.ListGamesHandler
	GetGame   corePort.GetGameHandler
}
