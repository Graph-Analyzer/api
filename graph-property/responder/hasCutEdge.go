package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type HasCutEdgeResponder interface {
	CreateSuccessfulResponse(dto domain.HasCutEdgeDTO)
}

type hasCutEdge struct {
	ctx *gin.Context
}

func NewHasCutEdgeResponder(ctx *gin.Context) HasCutEdgeResponder {
	return &hasCutEdge{
		ctx: ctx,
	}
}

type HasCutEdgeProjection struct {
	Status bool `json:"status" validate:"required" example:"true"`
}

func (r *hasCutEdge) CreateSuccessfulResponse(dto domain.HasCutEdgeDTO) {
	proj := HasCutEdgeProjection{
		Status: dto.Status,
	}

	r.ctx.JSON(200, proj)
}
