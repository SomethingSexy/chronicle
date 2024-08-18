package service

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/game"
	gameApplication "github.com/SomethingSexy/chronicle/internal/chronicle/core/application"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/go-chi/chi/v5"
)

func NewService() {
	game := gameApplication.NewApplication()

	service := ChronicleService{
		ChronicleApplication: game,
	}

	httpServer := http.NewHttpServer(service)

	httpServer.Start()
}

type ChronicleService struct {
	ChronicleApplication port.ChronicleApplication
}

func (c ChronicleService) Routes() []chi.Router {
	gameHttpServer := game.NewGameHttpServer(c.ChronicleApplication.Commands, port.GameQueries{})
	routes := gameHttpServer.Routes()
	return []chi.Router{routes}
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
