package action

import (
	"graph-analyzer/api/infrastructure/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type LivenessAction interface {
	_interface.Action
}

type liveness struct {
}

func NewLivenessAction() LivenessAction {
	return &liveness{}
}

// Invoke acts as the endpoint for liveness godoc
//
//	@Router			/live [get]
//	@Tags			infrastructure
//	@Summary		Get system liveness
//	@Description	Check if system is available (used in k8s)
//	@Produce		json
//	@Success		200
//	@Failure		500
func (a *liveness) Invoke(ctx *gin.Context) {
	res := responder.NewLivenessResponder(ctx)

	res.CreateSuccessfulResponse()
}
