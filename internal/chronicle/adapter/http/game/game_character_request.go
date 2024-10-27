package game

import (
	"io"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/port"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewGameCharacterRequest(body io.ReadCloser) (GameCharacterRequest, error) {
	// payload := new(onePayload)
	// At this point, in theory we don't know the type of game we are processing.
	// At the Http level, we just want to verify basic structure then, we will do a much
	// more indepth validation of the payload in the command.

	// TODO: Could we instead go back to jsonapi but make a sub property
	// that would jsut be a map interface?

	// err := json.NewDecoder(body).Decode(payload)
	// if err != nil {
	// 	return GameCharacterRequest{}, err
	// }

	// return GameCharacterRequest{
	// 	ID:   payload.Data.ID,
	// 	Data: payload.Data.Attributes,
	// }, nil

	var model GameCharacterRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

func NewGameCharacterResponse(gameCharacter domain.GameCharacter) GameCharacterRequest {
	return GameCharacterRequest{
		ID:          gameCharacter.GetId().String(),
		CharacterId: gameCharacter.GetCharacterId().String(),
		Character:   gameCharacter.GetData(),
		Type:        gameCharacter.GetType().String(),
	}
}

// A request to update a character tied to game through a
// world.  The schema for this character is based on the
// the type of game.
type GameCharacterRequest struct {
	ID          string                 `jsonapi:"primary,game-characters"`
	GameId      string                 `jsonapi:"attr,gameId"`
	CharacterId string                 `jsonapi:"attr,characterId"`
	Character   map[string]interface{} `jsonapi:"attr,character"`
	Type        string                 `jsonapi:"attr,type"`
}

// Having ToDomain here return the port instead of the struct
// directly.  We can't convert to the domain because we don't know the type
// yet.  Decide in the future if this is a good idea or not
func (g GameCharacterRequest) ToDomain() port.UpdateGameCharacter {
	return port.UpdateGameCharacter{
		GameId:      uuid.MustParse(g.GameId),
		CharacterId: uuid.MustParse(g.CharacterId),
		Character:   g.Character,
		Type:        domain.NewCharacterType(g.Type),
	}
}

// type onePayload struct {
// 	Data *node `json:"data"`
// }

// type node struct {
// 	Type          string                 `json:"type"`
// 	ID            string                 `json:"id,omitempty"`
// 	ClientID      string                 `json:"client-id,omitempty"`
// 	Attributes    map[string]interface{} `json:"attributes,omitempty"`
// 	Relationships map[string]interface{} `json:"relationships,omitempty"`
// }
