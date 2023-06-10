package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type DegreeCorrelationAction interface {
	_interface.Action
}

type degreeCorrelation struct {
	service domain.DegreeCorrelationService
}

func NewDegreeCorrelationAction(service domain.DegreeCorrelationService) DegreeCorrelationAction {
	return &degreeCorrelation{
		service: service,
	}
}

// Invoke acts as the endpoint for the degree correlation graph property godoc
//
//	@Router			/graph-property/degree-correlation [get]
//	@Tags			graph-properties
//	@Summary		Get degree correlation
//	@Description	Get degree correlation of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.DegreeCorrelationProjection
//	@Failure		500
func (a *degreeCorrelation) Invoke(ctx *gin.Context) {
	res := responder.NewDegreeCorrelationResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
