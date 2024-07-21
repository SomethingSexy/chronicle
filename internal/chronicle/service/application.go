package service

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/adapter"
	gameApplication "github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	gameApplication.GameCommands
}

type Queries struct {
	gameApplication.GameQueries
}

func NewService() {
	adapter.NewHttpServer()
}
