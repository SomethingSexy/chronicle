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

func NewGameHttpServer(commands port.GameCommands, queries port.GameQueries) GameHttpServer {
	return GameHttpServer{
		commands: commands,
		queries:  queries,
	}
}

type GameHttpServer struct {
	commands port.GameCommands
	queries  port.GameQueries
}

func (h GameHttpServer) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateGame)
	r.Get("/", h.ListGames)
	r.Get("/{gameId}", h.GetGame)
	// Not running this one as a relationship because
	// the character is technically already part of the game from the world
	// but this would be expanding on details for that character
	r.Post("/{gameId}/characters", h.UpdateGameCharacter)

	return r
}

func (h GameHttpServer) CreateGame(w http.ResponseWriter, r *http.Request) {
	data, err := NewGameRequest(r.Body)
	if err != nil {
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
		gameRequest := NewGameResponse((game))
		responses[i] = &gameRequest
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

	gameRequest := NewGameResponse(game)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, &gameRequest); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}

func (h GameHttpServer) UpdateGameCharacter(w http.ResponseWriter, r *http.Request) {
	data, err := NewGameCharacterRequest(r.Body)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.UpdateGameCharacter.Handle(r.Context(), data.ToDomain()); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}
