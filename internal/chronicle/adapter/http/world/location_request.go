package world

import (
	"errors"
	"io"
	"net/http"

	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NewLocationResponse(l domain.Location) LocationRequest {
	paths := make([]string, len(l.Path))
	for x, path := range l.Path {
		paths[x] = path.String()
	}

	return LocationRequest{
		ID:         l.LocationId.String(),
		LocationId: l.LocationId.String(),
		WorldId:    l.WorldId.String(),
		Name:       l.Name,
		Type:       l.Type,
		Path:       paths,
	}
}

func NewLocationRequest(body io.ReadCloser) (LocationRequest, error) {
	var model LocationRequest
	if err := jsonapi.UnmarshalPayload(body, &model); err != nil {
		return model, err
	}

	return model, nil
}

type LocationRequest struct {
	ID         string   `jsonapi:"primary,locations"`
	LocationId string   `jsonapi:"attr,locationId"`
	WorldId    string   `jsonapi:"attr,worldId"`
	Name       string   `jsonapi:"attr,name"`
	Type       string   `jsonapi:"attr,type"`
	Path       []string `jsonapi:"attr,path"`
}

func (a *LocationRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required name field")
	}

	if a.Type == "" {
		return errors.New("missing required type field")
	}

	if _, err := uuid.Parse(a.LocationId); err != nil {
		return errors.New("locationId must be valid UUID")
	}

	if _, err := uuid.Parse(a.WorldId); err != nil {
		return errors.New("worldId must be valid UUID")
	}

	for _, part := range a.Path {
		if _, err := uuid.Parse(part); err != nil {
			return errors.New("each parth within the path must be a valid UUID")
		}
	}

	return nil
}

func (a *LocationRequest) ToDomain() domain.Location {
	path := make([]uuid.UUID, len(a.Path))
	for i, part := range a.Path {
		path[i] = uuid.MustParse(part)
	}

	return domain.Location{
		LocationId: uuid.MustParse(a.LocationId),
		WorldId:    uuid.MustParse(a.WorldId),
		Name:       a.Name,
		Type:       a.Type,
		Path:       path,
	}
}
