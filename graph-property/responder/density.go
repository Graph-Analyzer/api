package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type DensityResponder interface {
	CreateSuccessfulResponse(dto domain.DensityDTO)
}

type density struct {
	ctx *gin.Context
}

func NewDensityResponder(ctx *gin.Context) DensityResponder {
	return &density{
		ctx: ctx,
	}
}

type DensityProjection struct {
	Value float64 `json:"value" validate:"required" example:"0.4684"`
}

func (r *density) CreateSuccessfulResponse(dto domain.DensityDTO) {
	proj := DensityProjection{
		Value: dto.Value,
	}

	r.ctx.JSON(200, proj)
}
