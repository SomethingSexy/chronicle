package adapter

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

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

type GameHttpServer struct {
}

func NewHttpServer() GameHttpServer {
	return GameHttpServer{}
}

func (h GameHttpServer) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", CreateGame)
	return r
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	data := &GameRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	// render.Render(w, r, NewGameResponse(article))
}

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
