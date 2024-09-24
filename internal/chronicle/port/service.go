package port

import (
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
)

type ChronicleApplication struct {
	Commands    ChronicleCommands
	Queries     GameQueries
	Persistence Persistence
}

type ChronicleCommands struct {
	CreateGame     corePort.CreateGameHander
	CreateWorld    corePort.CreateWorldHander
	CreateLocation corePort.CreateLocationHander
}

type GameQueries struct {
	GetGame       corePort.GetGameHandler
	GetWorld      corePort.GetWorldHandler
	ListGames     corePort.ListGamesHandler
	ListLocations corePort.ListLocationsHandler
}
