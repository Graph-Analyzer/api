package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type DiameterResponder interface {
	CreateSuccessfulResponse(dto domain.DiameterDTO)
}

type diameter struct {
	ctx *gin.Context
}

func NewDiameterResponder(ctx *gin.Context) DiameterResponder {
	return &diameter{
		ctx: ctx,
	}
}

type DiameterProjection struct {
	Diameter float64 `json:"diameter" validate:"required" example:"3.0"`
	Hops     int64   `json:"hops" validate:"required" example:"5"`
}

func (r *diameter) CreateSuccessfulResponse(dto domain.DiameterDTO) {
	proj := DiameterProjection{
		Diameter: dto.Diameter,
		Hops:     dto.Hops,
	}

	r.ctx.JSON(200, proj)
}
