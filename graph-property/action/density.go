package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type DensityAction interface {
	_interface.Action
}

type density struct {
	service domain.DensityService
}

func NewDensityAction(service domain.DensityService) DensityAction {
	return &density{
		service: service,
	}
}

// Invoke acts as the endpoint for the density graph property godoc
//
//	@Router			/graph-property/density [get]
//	@Tags			graph-properties
//	@Summary		Get density
//	@Description	Get density of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.DensityProjection
//	@Failure		500
func (a *density) Invoke(ctx *gin.Context) {
	res := responder.NewDensityResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
