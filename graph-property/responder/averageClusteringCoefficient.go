package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type AverageClusteringCoefficientResponder interface {
	CreateSuccessfulResponse(dto domain.AverageClusteringCoefficientDTO)
}

type averageClusteringCoefficient struct {
	ctx *gin.Context
}

func NewAverageClusteringCoefficientResponder(ctx *gin.Context) AverageClusteringCoefficientResponder {
	return &averageClusteringCoefficient{
		ctx: ctx,
	}
}

type AverageClusteringCoefficientProjection struct {
	Value float64 `json:"value" validate:"required" example:"0.5"`
}

func (r *averageClusteringCoefficient) CreateSuccessfulResponse(dto domain.AverageClusteringCoefficientDTO) {
	proj := AverageClusteringCoefficientProjection{
		Value: dto.Value,
	}

	r.ctx.JSON(200, proj)
}
