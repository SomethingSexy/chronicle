package world

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

func NewWorldHttpServer(commands port.WorldCommands, queries port.WorldQueries) GameHttpServer {
	return GameHttpServer{
		commands: commands,
		queries:  queries,
	}
}

type GameHttpServer struct {
	commands port.WorldCommands
	queries  port.WorldQueries
}

func (h GameHttpServer) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateWorld)
	r.Get("/{worldId}", h.GetWorld)

	// TODO: This should probably be relationships?
	r.Post("/{worldId}/locations", h.CreateLocation)
	r.Get("/{worldId}/locations", h.GetLocations)

	r.Post("/{worldId}/relationships/characters", h.AddWorldCharacter)
	// Using this format right now for patching a character that has been linked to a world
	// This might not be correct and instead it could include "world-characters"
	// TODO: I don't think we need this anymore
	// r.Patch("/{worldId}/characters/{characterId}", h.AddWorldCharacter)

	return r
}

func (h GameHttpServer) CreateWorld(w http.ResponseWriter, r *http.Request) {
	data, err := NewWorldRequest(r.Body)
	if err != nil {
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
	worldId := chi.URLParam(r, "worldId")

	world, err := h.queries.GetWorld.Handle(r.Context(), corePort.GetWorldQuery{
		WorldId: uuid.MustParse(worldId),
	})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	worldRquest := NewWorldResponse(world)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, &worldRquest); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}

func (h GameHttpServer) CreateLocation(w http.ResponseWriter, r *http.Request) {
	data, err := NewLocationRequest(r.Body)
	if err != nil {
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
	worldId := chi.URLParam(r, "worldId")

	locations, err := h.queries.ListLocations.Handle(r.Context(), corePort.LocationsQuery{
		WorldId: uuid.MustParse(worldId),
	})
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	responses := make([]*LocationRequest, len(locations))
	for i, location := range locations {
		locationRequest := NewLocationResponse(location)
		responses[i] = &locationRequest
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(w, responses); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
}

func (h GameHttpServer) AddWorldCharacter(w http.ResponseWriter, r *http.Request) {
	worldId := chi.URLParam(r, "worldId")
	data, err := NewAddWorldCharacterRequest(r.Body)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.AddWorldCharacter.Handle(r.Context(), corePort.AddWorldCharacter{
		// TODO: Bad check for error
		WorldId:     uuid.MustParse(worldId),
		CharacterId: data.ToDomain(),
	}); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}
