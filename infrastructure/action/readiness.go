package action

import (
	"graph-analyzer/api/infrastructure/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type ReadinessAction interface {
	_interface.Action
}

type readiness struct {
}

func NewReadinessAction() ReadinessAction {
	return &readiness{}
}

// Invoke acts as the endpoint for readiness godoc
//
//	@Router			/ready [get]
//	@Tags			infrastructure
//	@Summary		Get system readiness
//	@Description	Check if system is ready (used in k8s)
//	@Produce		json
//	@Success		200
//	@Failure		500
func (a *readiness) Invoke(ctx *gin.Context) {
	res := responder.NewReadinessResponder(ctx)

	res.CreateSuccessfulResponse()
}
