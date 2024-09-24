package service

import (
	"context"
	"fmt"
	"os"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/game"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewService() {
	// TODO: This shold probably come from the adapter instead
	psqlConnStr := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), psqlConnStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	q := repository.New(db)

	// TODO: We will need to create these individally as we add and merge together
	persistence := port.Persistence{
		Game:      query.NewGameQuery(q),
		Character: query.NewCharacterQuery(q),
	}

	app := application.NewApplication(persistence)

	service := ChronicleService{
		ChronicleApplication: app,
	}

	httpServer := http.NewHttpServer(service)

	httpServer.Start()
}

type ChronicleService struct {
	ChronicleApplication port.ChronicleApplication
}

func (c ChronicleService) Routes() []chi.Router {
	gameHttpServer := game.NewGameHttpServer(c.ChronicleApplication.Commands, c.ChronicleApplication.Queries)
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
