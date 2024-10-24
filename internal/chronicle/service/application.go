package service

import (
	"context"
	"fmt"
	"os"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/character"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/game"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/http/world"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/query"
	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/application"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewService() error {
	// TODO: This shold probably come from the adapter instead
	psqlConnStr := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), psqlConnStr)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	q := repository.New(db)

	persistence := port.Persistence{
		Game:      query.NewGameQuery(q),
		Character: query.NewCharacterQuery(q),
		World:     query.NewWorldQuery(q),
	}

	app := application.NewApplication(persistence)

	service := ChronicleService{
		ChronicleApplication: app,
	}

	httpServer := http.NewHttpServer(service)

	return httpServer.Start()
}

type ChronicleService struct {
	ChronicleApplication port.ChronicleApplication
}

func (c ChronicleService) Routes() map[string][]chi.Router {
	// TODO: This could probably get created in the adapter itself, less this package needs to import
	gameHttpServer := game.NewGameHttpServer(c.ChronicleApplication.Commands.GameCommands, c.ChronicleApplication.Queries.GameQueries)
	gameRoutes := gameHttpServer.Routes()

	characterHttpServer := character.NewCharacterHttpServer(c.ChronicleApplication.Commands.CharacterCommands, c.ChronicleApplication.Queries.CharacterQueries)
	characterRoutes := characterHttpServer.Routes()

	worldHttpServer := world.NewWorldHttpServer(c.ChronicleApplication.Commands.WorldCommands, c.ChronicleApplication.Queries.WorldQueries)
	worldRoutes := worldHttpServer.Routes()

	return map[string][]chi.Router{
		"Games":      {gameRoutes},
		"Characters": {characterRoutes},
		"Worlds":     {worldRoutes},
	}
}
