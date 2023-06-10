package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type HasCutEdgeAction interface {
	_interface.Action
}

type hasCutEdge struct {
	service domain.HasCutEdgeService
}

func NewHasCutEdgeAction(service domain.HasCutEdgeService) HasCutEdgeAction {
	return &hasCutEdge{
		service: service,
	}
}

// Invoke acts as the endpoint for the has cut edge graph property godoc
//
//	@Router			/graph-property/has-cut-edge [get]
//	@Tags			graph-properties
//	@Summary		Get has cut edge
//	@Description	Get if graph has a cut edge. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.HasCutEdgeProjection
//	@Failure		500
func (a *hasCutEdge) Invoke(ctx *gin.Context) {
	res := responder.NewHasCutEdgeResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
