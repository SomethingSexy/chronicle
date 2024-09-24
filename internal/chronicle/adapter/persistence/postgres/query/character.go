package query

import (
	"context"
	"time"

	"github.com/SomethingSexy/chronicle/internal/chronicle/adapter/persistence/postgres/sqlc/repository"
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewCharacterQuery(queries *repository.Queries) CharacterQuery {
	return CharacterQuery{
		Queries: queries,
	}
}

type CharacterQuery struct {
	Queries *repository.Queries
}

func (c CharacterQuery) CreateCharacter(ctx context.Context, character domain.Character) (domain.Character, error) {
	ts := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}
	_, err := c.Queries.CreateCharacter(ctx, repository.CreateCharacterParams{
		CharacterID: character.CharacterId,
		Name:        character.Name,
		Description: pgtype.Text{
			String: character.Description,
			Valid:  true,
		},
		CreatedAt: ts,
		UpdatedAt: ts,
	})
	if err != nil {
		return domain.Character{}, err
	}

	return character, nil
}
