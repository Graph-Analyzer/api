package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type ConnectedResponder interface {
	CreateSuccessfulResponse(dto domain.ConnectedDTO)
}

type connected struct {
	ctx *gin.Context
}

func NewConnectedResponder(ctx *gin.Context) ConnectedResponder {
	return &connected{
		ctx: ctx,
	}
}

type ConnectedProjection struct {
	Status bool `json:"status" validate:"required" example:"true"`
}

func (r *connected) CreateSuccessfulResponse(dto domain.ConnectedDTO) {
	proj := ConnectedProjection{
		Status: dto.Status,
	}

	r.ctx.JSON(200, proj)
}
