package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type HasCutVertexResponder interface {
	CreateSuccessfulResponse(dto domain.HasCutVertexDTO)
}

type hasCutVertex struct {
	ctx *gin.Context
}

func NewHasCutVertexResponder(ctx *gin.Context) HasCutVertexResponder {
	return &hasCutVertex{
		ctx: ctx,
	}
}

type HasCutVertexProjection struct {
	Status bool `json:"status" validate:"required" example:"true"`
}

func (r *hasCutVertex) CreateSuccessfulResponse(dto domain.HasCutVertexDTO) {
	proj := HasCutVertexProjection{
		Status: dto.Status,
	}

	r.ctx.JSON(200, proj)
}
