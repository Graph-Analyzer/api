package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type DiameterAction interface {
	_interface.Action
}

type diameter struct {
	service domain.DiameterService
}

func NewDiameterAction(service domain.DiameterService) DiameterAction {
	return &diameter{
		service: service,
	}
}

// Invoke acts as the endpoint for the diameter graph property godoc
//
//	@Router			/graph-property/diameter [get]
//	@Tags			graph-properties
//	@Summary		Get diameter
//	@Description	Get diameter of the graph. Based on directed and weighted multi graph.
//	@Produce		json
//	@Success		200	{object}	responder.DiameterProjection
//	@Failure		500
func (a *diameter) Invoke(ctx *gin.Context) {
	res := responder.NewDiameterResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
