package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type RobustnessResponder interface {
	CreateSuccessfulResponse(dto domain.RobustnessDTO)
}

type robustness struct {
	ctx *gin.Context
}

func NewRobustnessResponder(ctx *gin.Context) RobustnessResponder {
	return &robustness{
		ctx: ctx,
	}
}

type RobustnessProjection struct {
	Nodes      int64   `json:"nodes" validate:"required" example:"3"`
	Percentage float64 `json:"percentage" validate:"required" example:"33.446"`
}

func (r *robustness) CreateSuccessfulResponse(dto domain.RobustnessDTO) {
	proj := RobustnessProjection{
		Nodes:      dto.Nodes,
		Percentage: dto.Percentage,
	}

	r.ctx.JSON(200, proj)
}
