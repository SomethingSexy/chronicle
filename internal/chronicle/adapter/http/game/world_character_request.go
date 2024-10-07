package game

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type WorldCharacterRequest struct {
	ID string `jsonapi:"primary,characters"`
}

func (a *WorldCharacterRequest) Bind(r *http.Request) error {
	if _, err := uuid.Parse(a.ID); err != nil {
		return errors.New("id must be valid UUID")
	}

	return nil
}

func (a *WorldCharacterRequest) ToDomain() uuid.UUID {
	return uuid.MustParse(a.ID)
}
