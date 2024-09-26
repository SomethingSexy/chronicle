package port

import (
	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
)

type ChronicleApplication struct {
	Commands    ChronicleCommands
	Queries     ChronicleQueries
	Persistence Persistence
}

type ChronicleCommands struct {
	GameCommands
	CharacterCommands
}

type ChronicleQueries struct {
	GameQueries
	CharacterQueries
}

type GameCommands struct {
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

type CharacterCommands struct {
	CreateCharacter corePort.CreateCharacterHandler
}

type CharacterQueries struct{}
