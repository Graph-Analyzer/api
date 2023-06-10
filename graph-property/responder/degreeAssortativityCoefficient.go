package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type DegreeAssortativityCoefficientResponder interface {
	CreateSuccessfulResponse(dto domain.DegreeAssortativityCoefficientDTO)
}

type degreeAssortativityCoefficient struct {
	ctx *gin.Context
}

func NewDegreeAssortativityCoefficientResponder(ctx *gin.Context) DegreeAssortativityCoefficientResponder {
	return &degreeAssortativityCoefficient{
		ctx: ctx,
	}
}

type DegreeAssortativityCoefficientProjection struct {
	Value float64 `json:"value" validate:"required" example:"0.5"`
}

func (r *degreeAssortativityCoefficient) CreateSuccessfulResponse(dto domain.DegreeAssortativityCoefficientDTO) {
	proj := DegreeAssortativityCoefficientProjection{
		Value: dto.Value,
	}

	r.ctx.JSON(200, proj)
}
