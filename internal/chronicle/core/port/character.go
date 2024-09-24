package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
)

type CreateCharacter struct {
	Character domain.Character
}

type CreateCharacterHander common.CommandHandler[CreateCharacter]
