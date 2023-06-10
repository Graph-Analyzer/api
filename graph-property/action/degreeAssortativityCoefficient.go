package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type DegreeAssortativityCoefficientAction interface {
	_interface.Action
}

type degreeAssortativityCoefficient struct {
	service domain.DegreeAssortativityCoefficientService
}

func NewDegreeAssortativityCoefficientAction(service domain.DegreeAssortativityCoefficientService) DegreeAssortativityCoefficientAction {
	return &degreeAssortativityCoefficient{
		service: service,
	}
}

// Invoke acts as the endpoint for the degree assortativity coefficient graph property godoc
//
//	@Router			/graph-property/degree-assortativity-coefficient [get]
//	@Tags			graph-properties
//	@Summary		Get degree assortativity coefficient
//	@Description	Get degree assortativity coefficient of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.DegreeAssortativityCoefficientProjection
//	@Failure		500
func (a *degreeAssortativityCoefficient) Invoke(ctx *gin.Context) {
	res := responder.NewDegreeAssortativityCoefficientResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
