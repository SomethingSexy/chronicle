package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
)

type CreateWorld struct {
	World domain.World
}

type CreateWorldHander common.CommandHandler[CreateWorld]
