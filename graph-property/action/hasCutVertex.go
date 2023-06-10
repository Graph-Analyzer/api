package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type HasCutVertexAction interface {
	_interface.Action
}

type hasCutVertex struct {
	service domain.HasCutVertexService
}

func NewHasCutVertexAction(service domain.HasCutVertexService) HasCutVertexAction {
	return &hasCutVertex{
		service: service,
	}
}

// Invoke acts as the endpoint for the has cut vertex graph property godoc
//
//	@Router			/graph-property/has-cut-vertex [get]
//	@Tags			graph-properties
//	@Summary		Get has cut vertex
//	@Description	Get if graph has a cut vertex. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.HasCutVertexProjection
//	@Failure		500
func (a *hasCutVertex) Invoke(ctx *gin.Context) {
	res := responder.NewHasCutVertexResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
