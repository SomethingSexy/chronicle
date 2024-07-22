package port

import "github.com/go-chi/chi/v5"

// This should represent an overall service application
type Service interface {
	Routes() []chi.Router
}
