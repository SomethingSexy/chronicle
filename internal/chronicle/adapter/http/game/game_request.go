package game

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/uuid"
)

type GameRequest struct {
	ID     string          `jsonapi:"primary,games"`
	GameId string          `jsonapi:"attr,gameId"`
	Name   string          `jsonapi:"attr,name"`
	Type   string          `jsonapi:"attr,type"`
	Worlds []*WorldRequest `jsonapi:"relation,worlds"`
	// ID           int    `jsonapi:"primary,blogs"`
	// Title        string `jsonapi:"attr,title"`
	// Posts         []*Post   `jsonapi:"relation,posts"`
	// CurrentPost   *Post     `jsonapi:"relation,current_post"`
	// CurrentPostID int       `jsonapi:"attr,current_post_id"`
	// CreatedAt     time.Time `jsonapi:"attr,created_at"`
	// ViewCount     int       `jsonapi:"attr,view_count"`
}

func (a *GameRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required game fields")
	}

	if a.Type == "" {
		return errors.New("missing required game fields")
	}

	if _, err := uuid.Parse(a.GameId); err != nil {
		return errors.New("gameId must be valid UUID")
	}

	return nil
}

func (a *GameRequest) ToDomain() domain.Game {
	return domain.Game{
		GameId: uuid.MustParse(a.GameId),
		Name:   a.Name,
		Type:   a.Type,
	}
}
