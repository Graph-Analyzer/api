package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type DegreeDistributionResponder interface {
	CreateSuccessfulResponse(dto domain.DegreeDistributionDTO)
}

type degreeDistribution struct {
	ctx *gin.Context
}

func NewDegreeDistributionResponder(ctx *gin.Context) DegreeDistributionResponder {
	return &degreeDistribution{
		ctx: ctx,
	}
}

type DegreeDistributionProjection struct {
	Values []DegreeDistributionValueProjection `json:"values" validate:"required"`
}

type DegreeDistributionValueProjection struct {
	Degree int64 `json:"degree" validate:"required" example:"4"`
	Amount int64 `json:"amount" validate:"required" example:"4"`
}

func (r *degreeDistribution) CreateSuccessfulResponse(dto domain.DegreeDistributionDTO) {
	values := make([]DegreeDistributionValueProjection, 0)

	for _, value := range dto.Values {
		values = append(values, DegreeDistributionValueProjection{
			Degree: value.Degree,
			Amount: value.Amount,
		})
	}

	proj := DegreeDistributionProjection{
		Values: values,
	}

	r.ctx.JSON(200, proj)
}
