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

type HttpServer struct {
	app common.Service
}

func NewHttpServer(application common.Service) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Start() {
	// This should be responsible for running and starting the http service but it should be given the routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	render.Decode = DefaultDecoder

	// TODO: Given the application, this should mount all of the route handlers
	r.Mount("/games", h.app.Routes()[0])

	http.ListenAndServe(":3000", r)
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
