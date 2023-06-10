package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type RobustnessAction interface {
	_interface.Action
}

type robustness struct {
	service domain.RobustnessService
}

func NewRobustnessAction(service domain.RobustnessService) RobustnessAction {
	return &robustness{
		service: service,
	}
}

// Invoke acts as the endpoint for the robustness graph property godoc
//
//	@Router			/graph-property/robustness [get]
//	@Tags			graph-properties
//	@Summary		Get robustness
//	@Description	Get robustness of the graph for a high-degree attack. Based on and unweighted simple graph.
//	@Produce		json
//	@Success		200	{object}	responder.RobustnessProjection
//	@Failure		500
func (a *robustness) Invoke(ctx *gin.Context) {
	res := responder.NewRobustnessResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
