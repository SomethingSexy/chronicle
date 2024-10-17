package character

import (
	"net/http"

	corePort "github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/SomethingSexy/chronicle/internal/chronicle/port"
	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewCharacterHttpServer(commands port.CharacterCommands, queries port.CharacterQueries) CharacterHttpServer {
	return CharacterHttpServer{
		commands: commands,
		queries:  queries,
	}
}

type CharacterHttpServer struct {
	commands port.CharacterCommands
	queries  port.CharacterQueries
}

func (h CharacterHttpServer) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateCharacter)

	return r
}

func (h CharacterHttpServer) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	data, err := NewCharacterRequest(r.Body)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := h.commands.CreateCharacter.Handle(r.Context(), corePort.CreateCharacter{
		Character: data.ToDomain(),
	}); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
}
