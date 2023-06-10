package responder

import (
	"github.com/gin-gonic/gin"
)

type ReadinessResponder interface {
	CreateSuccessfulResponse()
}

type readiness struct {
	ctx *gin.Context
}

func NewReadinessResponder(ctx *gin.Context) ReadinessResponder {
	return &readiness{
		ctx: ctx,
	}
}

func (r *readiness) CreateSuccessfulResponse() {
	r.ctx.Status(200)
}
