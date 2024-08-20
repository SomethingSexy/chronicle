package game

import (
	"net/http"

	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewGameHttpServer(commands port.ChronicleCommands, queries port.GameQueries) GameHttpServer {
	return GameHttpServer{
		commands: commands,
		queries:  queries,
	}
}

type GameHttpServer struct {
	commands port.ChronicleCommands
	queries  port.GameQueries
}

func (h GameHttpServer) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateGame)
	r.Get("/", h.ListGames)
	r.Get("/{gameId}", h.GetGame)
	return r
}

func (h GameHttpServer) CreateGame(w http.ResponseWriter, r *http.Request) {
	data := &GameRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.CreateGame.Handle(r.Context(), corePort.CreateGame{
		Game: data.ToDomain(),
	}); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}

func (h GameHttpServer) ListGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.queries.ListGames.Handle(r.Context(), corePort.AllGamesQuery{})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	responses := make([]*GameRequest, len(games))
	for i, game := range games {
		responses[i] = &GameRequest{
			ID:     game.GameId.String(),
			GameId: game.GameId.String(),
			Name:   game.Name,
			Type:   game.Type,
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, responses); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}

func (h GameHttpServer) GetGame(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "gameId")

	game, err := h.queries.GetGame.Handle(r.Context(), corePort.GetGameQuery{
		// TODO: BAD check for error
		GameId: uuid.MustParse(gameId),
	})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	response := &GameRequest{
		ID:     game.GameId.String(),
		GameId: game.GameId.String(),
		Name:   game.Name,
		Type:   game.Type,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, response); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}
