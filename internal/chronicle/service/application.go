package service

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter"
	gameApplication "github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application"
	gamePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/game/port"
	"github.com/go-chi/chi/v5"
)

func NewService() {
	game := gameApplication.NewApplication()

	service := ChronicleService{
		GameApplication: game,
	}

	httpServer := adapter.NewHttpServer(service)

	httpServer.Start()
}

type ChronicleService struct {
	GameApplication gamePort.GameApplication
}

func (c ChronicleService) Routes() []chi.Router {
	gameHttpServer := c.GameApplication.Server.Routes()
	return []chi.Router{gameHttpServer}
}

// type Application struct {
// 	Commands Commands
// 	Queries  Queries
// 	Routes   []chi.Router
// }

// type Commands struct {
// 	gameApplication.GameCommands
// }

// type Queries struct {
// 	gameApplication.GameQueries
// }
