package responder

import (
	"github.com/gin-gonic/gin"
)

type LivenessResponder interface {
	CreateSuccessfulResponse()
}

type liveness struct {
	ctx *gin.Context
}

func NewLivenessResponder(ctx *gin.Context) LivenessResponder {
	return &liveness{
		ctx: ctx,
	}
}

func (r *liveness) CreateSuccessfulResponse() {
	r.ctx.Status(200)
}
