package world

import (
	"errors"
	"io"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewAddWorldCharacterRequest(body io.ReadCloser) (AddWorldCharacterRequest, error) {
	var model AddWorldCharacterRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

// Rquest typed used for add a character to a world via a relationship
type AddWorldCharacterRequest struct {
	ID string `jsonapi:"primary,characters"`
}

func (a *AddWorldCharacterRequest) Bind(r *http.Request) error {
	if _, err := uuid.Parse(a.ID); err != nil {
		return errors.New("id must be valid UUID")
	}

	return nil
}

func (a *AddWorldCharacterRequest) ToDomain() uuid.UUID {
	return uuid.MustParse(a.ID)
}

// func NewWorldCharacterRequest(body io.ReadCloser) (WorldCharacterRequest, error) {
// 	var model WorldCharacterRequest
// 	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
// 		return model, err
// 	}

// 	return model, nil
// }

// Request type used for updating the attributes for a
// character that has already been added to a world
// type WorldCharacterRequest struct {
// 	ID   string `jsonapi:"primary,world-characters"`
// 	Type string `jsonapi:"attr,type"`
// }

// func (a *WorldCharacterRequest) Bind(r *http.Request) error {
// 	if _, err := uuid.Parse(a.ID); err != nil {
// 		return errors.New("id must be valid UUID")
// 	}

// 	return nil
// }

// func (a *WorldCharacterRequest) ToDomain() domain.WorldCharacter {
// 	return domain.WorldCharacter{
// 		Type: domain.NewCharacterType(a.Type),
// 	}
// }
