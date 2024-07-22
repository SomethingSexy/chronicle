package adapter

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/application/command"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/port"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewHttpServer(commands port.GameCommands, queries port.GameQueries) GameHttpServer {
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
	return r
}

func (h GameHttpServer) CreateGame(w http.ResponseWriter, r *http.Request) {
	data := &GameRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if err := h.commands.CreateGame.Handle(r.Context(), command.CreateGame{
		Game: *data.Game,
	}); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	// render.Render(w, r, NewGameResponse(article))
}

// TODO: Do something with this error shit
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type GameRequest struct {
	*domain.Game
}

func (a *GameRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if a.Game == nil {
		return errors.New("missing required game fields")
	}

	return nil
}
