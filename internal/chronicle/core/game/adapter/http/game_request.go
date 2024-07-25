package http

import (
	"errors"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/game/domain"
)

type GameRequest struct {
	Name string `jsonapi:"attr,name"`
	Type string `jsonapi:"attr,type"`
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

	return nil
}

func (a *GameRequest) ToDomain() domain.Game {
	return domain.Game{
		Type: a.Type,
		Name: a.Name,
	}
}
