package port

import "github.com/go-chi/chi/v5"

type HttpServer interface {
	Routes() chi.Router
}
