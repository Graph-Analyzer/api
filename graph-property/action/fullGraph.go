package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type FullGraphAction interface {
	_interface.Action
}

type fullGraph struct {
	service domain.FullGraphService
}

func NewFullGraphAction(service domain.FullGraphService) FullGraphAction {
	return &fullGraph{
		service: service,
	}
}

// Invoke acts as the endpoint for the full graph godoc
//
//	@Router			/graph-property/full-graph [get]
//	@Tags			graph-properties
//	@Summary		Get full graph
//	@Description	Get full network graph (nodes, edges)
//	@Produce		json
//	@Success		200	{object}	responder.FullGraphProjection
//	@Failure		500
func (a *fullGraph) Invoke(ctx *gin.Context) {
	res := responder.NewFullGraphResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
