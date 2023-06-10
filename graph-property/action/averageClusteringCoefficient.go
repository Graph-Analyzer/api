package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type AverageClusteringCoefficientAction interface {
	_interface.Action
}

type averageClusteringCoefficient struct {
	service domain.AverageClusteringCoefficientService
}

func NewAverageClusteringCoefficientAction(service domain.AverageClusteringCoefficientService) AverageClusteringCoefficientAction {
	return &averageClusteringCoefficient{
		service: service,
	}
}

// Invoke acts as the endpoint for the average clustering coefficient graph property godoc
//
//	@Router			/graph-property/average-clustering-coefficient [get]
//	@Tags			graph-properties
//	@Summary		Get average clustering coefficient
//	@Description	Get average clustering coefficient of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.AverageClusteringCoefficientProjection
//	@Failure		500
func (a *averageClusteringCoefficient) Invoke(ctx *gin.Context) {
	res := responder.NewAverageClusteringCoefficientResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
