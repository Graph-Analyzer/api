package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type ConnectedAction interface {
	_interface.Action
}

type connected struct {
	service domain.ConnectedService
}

func NewConnectedAction(service domain.ConnectedService) ConnectedAction {
	return &connected{
		service: service,
	}
}

// Invoke acts as the endpoint for the connected graph property godoc
//
//	@Router			/graph-property/connected [get]
//	@Tags			graph-properties
//	@Summary		Get connected status
//	@Description	Get connected status of the graph. Based on undirected and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.ConnectedProjection
//	@Failure		500
func (a *connected) Invoke(ctx *gin.Context) {
	res := responder.NewConnectedResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
