package http

import (
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/jsonapi"
)

var ContentTypeJsonApi = "application/vnd.api+json"

type Parser interface {
	Parse()
}

type HttpServer struct {
	app common.Service
}

// TODO: Does it make more sense that we pass in the core interfaces here
// and keep the route logic internal to this package, instead of passing it in?
func NewHttpServer(application common.Service) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Start() error {
	// This should be responsible for running and starting the http service but it should be given the routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	render.Decode = DefaultDecoder

	// TODO: Given the application, this should mount all of the route handlers
	r.Mount("/games", h.app.Routes()["Games"][0])
	r.Mount("/characters", h.app.Routes()["Characters"][0])
	r.Mount("/worlds", h.app.Routes()["Worlds"][0])

	return http.ListenAndServe(":3000", r)
}

func DefaultDecoder(r *http.Request, v interface{}) error {
	var err error

	switch r.Header.Get("Content-Type") {
	case ContentTypeJsonApi:
		if err := jsonapi.UnmarshalPayload(r.Body, v); err != nil {
			return err
		}
	default:
		err = render.DefaultDecoder(r, v)
	}

	return err
}
