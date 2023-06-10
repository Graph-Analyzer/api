package responder

import (
	"graph-analyzer/api/upload/domain"

	"github.com/gin-gonic/gin"
)

type UploadStatusResponder interface {
	CreateSuccessfulResponse(dto domain.UploadStatusDTO)
}

type uploadStatus struct {
	ctx *gin.Context
}

func NewUploadStatusResponder(ctx *gin.Context) UploadStatusResponder {
	return &uploadStatus{
		ctx: ctx,
	}
}

type UploadStatusProjection struct {
	Healthy bool `json:"healthy" validate:"required" example:"true"`
}

func (r *uploadStatus) CreateSuccessfulResponse(dto domain.UploadStatusDTO) {
	proj := UploadStatusProjection{
		Healthy: dto.Healthy,
	}

	r.ctx.JSON(200, proj)
}
