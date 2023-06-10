package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type DegreeDistributionAction interface {
	_interface.Action
}

type degreeDistribution struct {
	service domain.DegreeDistributionService
}

func NewDegreeDistributionAction(service domain.DegreeDistributionService) DegreeDistributionAction {
	return &degreeDistribution{
		service: service,
	}
}

// Invoke acts as the endpoint for the degree distribution graph property godoc
//
//	@Router			/graph-property/degree-distribution [get]
//	@Tags			graph-properties
//	@Summary		Get degree distribution
//	@Description	Get degree distribution of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.DegreeDistributionProjection
//	@Failure		500
func (a *degreeDistribution) Invoke(ctx *gin.Context) {
	res := responder.NewDegreeDistributionResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
