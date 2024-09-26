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

	// TODO: This is obviously going to get huge
	// we need to decide how to organize these route handlers
	// Either at this level or a higher level to apply
	// relationships to root routes.
	r.Post("/{gameId}/worlds", h.CreateWorld)
	r.Get("/{gameId}/worlds/{worldId}", h.GetWorld)
	r.Post("/{gameId}/worlds/{worldId}/locations", h.CreateLocation)
	r.Get("/{gameId}/worlds/{worldId}/locations", h.GetLocations)

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
		worlds := make([]*WorldRequest, len(game.Worlds))

		for x, world := range game.Worlds {
			worlds[x] = &WorldRequest{
				ID:      world.WorldId.String(),
				WorldId: world.WorldId.String(),
				GameId:  world.GameId.String(),
				Name:    world.Name,
			}
		}
		responses[i] = &GameRequest{
			ID:     game.GameId.String(),
			GameId: game.GameId.String(),
			Name:   game.Name,
			Type:   game.Type,
			Worlds: worlds,
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

func (h GameHttpServer) CreateWorld(w http.ResponseWriter, r *http.Request) {
	data := &WorldRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.CreateWorld.Handle(r.Context(), corePort.CreateWorld{
		World: data.ToDomain(),
	}); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}

func (h GameHttpServer) GetWorld(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "gameId")
	worldId := chi.URLParam(r, "worldId")

	world, err := h.queries.GetWorld.Handle(r.Context(), corePort.GetWorldQuery{
		// TODO: BAD check for error
		GameId:  uuid.MustParse(gameId),
		WorldId: uuid.MustParse(worldId),
	})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	locations := make([]*LocationRequest, len(world.Locations))
	for i, location := range world.Locations {
		paths := make([]string, len(location.Path))
		for x, path := range location.Path {
			paths[x] = path.String()
		}

		locations[i] = &LocationRequest{
			ID:         location.LocationId.String(),
			LocationId: location.LocationId.String(),
			WorldId:    location.WorldId.String(),
			Name:       location.Name,
			Type:       location.Type,
			Path:       paths,
		}
	}

	response := &WorldRequest{
		ID:        world.WorldId.String(),
		WorldId:   world.WorldId.String(),
		GameId:    world.GameId.String(),
		Name:      world.Name,
		Locations: locations,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, response); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}

func (h GameHttpServer) CreateLocation(w http.ResponseWriter, r *http.Request) {
	data := &LocationRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.CreateLocation.Handle(r.Context(), corePort.CreateLocation{
		Location: data.ToDomain(),
	}); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}

func (h GameHttpServer) GetLocations(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "gameId")
	worldId := chi.URLParam(r, "worldId")

	locations, err := h.queries.ListLocations.Handle(r.Context(), corePort.LocationsQuery{
		// TODO: BAD check for error
		GameId:  uuid.MustParse(gameId),
		WorldId: uuid.MustParse(worldId),
	})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	responses := make([]*LocationRequest, len(locations))

	for i, location := range locations {
		paths := make([]string, len(location.Path))
		for x, path := range location.Path {
			paths[x] = path.String()
		}

		responses[i] = &LocationRequest{
			ID:         location.LocationId.String(),
			LocationId: location.LocationId.String(),
			WorldId:    location.WorldId.String(),
			Name:       location.Name,
			Type:       location.Type,
			Path:       paths,
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, responses); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}
