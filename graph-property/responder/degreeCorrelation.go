package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type DegreeCorrelationResponder interface {
	CreateSuccessfulResponse(dto domain.DegreeCorrelationDTO)
}

type degreeCorrelation struct {
	ctx *gin.Context
}

func NewDegreeCorrelationResponder(ctx *gin.Context) DegreeCorrelationResponder {
	return &degreeCorrelation{
		ctx: ctx,
	}
}

type DegreeCorrelationProjection struct {
	Values []DegreeCorrelationValueProjection `json:"values" validate:"required"`
}

type DegreeCorrelationValueProjection struct {
	Degree  int64   `json:"degree" validate:"required" example:"4"`
	Average float64 `json:"average" validate:"required" example:"1.66"`
}

func (r *degreeCorrelation) CreateSuccessfulResponse(dto domain.DegreeCorrelationDTO) {
	values := make([]DegreeCorrelationValueProjection, 0)

	for _, value := range dto.Values {
		values = append(values, DegreeCorrelationValueProjection{
			Degree:  value.Degree,
			Average: value.Average,
		})
	}

	proj := DegreeCorrelationProjection{
		Values: values,
	}

	r.ctx.JSON(200, proj)
}
