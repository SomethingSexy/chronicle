package common

import "github.com/go-chi/chi/v5"

// Represents an overall rest http based service
type RestService interface {
	Routes() map[string][]chi.Router
	Port() string
}
