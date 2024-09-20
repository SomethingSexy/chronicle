package port

import (
	"github.com/SomethingSexy/chronicle/internal/chronicle/core/domain"
	"github.com/SomethingSexy/chronicle/internal/common"
)

type CreateLocation struct {
	Location domain.Location
}

type CreateLocationHander common.CommandHandler[CreateLocation]
