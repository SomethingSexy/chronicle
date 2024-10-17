package character

import (
	"errors"
	"io"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewCharacterResponse(c domain.Character) CharacterRequest {
	return CharacterRequest{
		ID:          c.CharacterId.String(),
		CharacterId: c.CharacterId.String(),
		Name:        c.Name,
		Description: c.Description,
	}
}

func NewCharacterRequest(body io.ReadCloser) (CharacterRequest, error) {
	var model CharacterRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

type CharacterRequest struct {
	ID          string `jsonapi:"primary,characters"`
	CharacterId string `jsonapi:"attr,characterId"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
}

func (c *CharacterRequest) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required name fields")
	}

	if _, err := uuid.Parse(c.CharacterId); err != nil {
		return errors.New("characterId must be valid UUID")
	}

	return nil
}

func (c *CharacterRequest) ToDomain() domain.Character {
	return domain.Character{
		CharacterId: uuid.MustParse(c.CharacterId),
		Name:        c.Name,
		Description: c.Description,
	}
}
