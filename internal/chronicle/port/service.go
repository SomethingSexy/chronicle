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
	WorldCommands
}

type ChronicleQueries struct {
	GameQueries
	CharacterQueries
	WorldQueries
}

type GameCommands struct {
	CreateGame        corePort.CreateGameHander
	CreateLocation    corePort.CreateLocationHander
	AddWorldCharacter corePort.AddWorldCharacterHandler
}

type GameQueries struct {
	GetGame   corePort.GetGameHandler
	ListGames corePort.ListGamesHandler
}

type WorldCommands struct {
	CreateWorld       corePort.CreateWorldHander
	CreateLocation    corePort.CreateLocationHander
	AddWorldCharacter corePort.AddWorldCharacterHandler
}

type WorldQueries struct {
	GetGame       corePort.GetGameHandler
	GetWorld      corePort.GetWorldHandler
	ListLocations corePort.ListLocationsHandler
}

type CharacterCommands struct {
	CreateCharacter corePort.CreateCharacterHandler
}

type CharacterQueries struct{}
