package port

import (
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	app service.Application
}

func NewHttpServer(application service.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Start() {
	// This should be responsible for running and starting the http service but it should be given the routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// TODO: Given the application, this should mount all of the route handlers

	http.ListenAndServe(":3000", r)
}
