package responder

import (
	"graph-analyzer/api/upload/domain"

	"github.com/gin-gonic/gin"
)

type UploadResponder interface {
	CreateSuccessfulResponse(dto domain.UploadDTO)
	CreateBadRequestErrorResponse(error string)
	CreateInternalErrorResponse(error string)
}

type upload struct {
	ctx *gin.Context
}

func NewUploadResponder(ctx *gin.Context) UploadResponder {
	return &upload{
		ctx: ctx,
	}
}

type UploadProjection struct {
	Status bool `json:"status" validate:"required" example:"true"`
}

type UploadErrorProjection struct {
	Error string `json:"error" validate:"error" example:"Filesize to large"`
}

func (r *upload) CreateSuccessfulResponse(dto domain.UploadDTO) {
	proj := UploadProjection{
		Status: dto.Status,
	}

	r.ctx.JSON(200, proj)
}

func (r *upload) CreateBadRequestErrorResponse(error string) {
	proj := UploadErrorProjection{
		Error: error,
	}

	r.ctx.JSON(400, proj)
}

func (r *upload) CreateInternalErrorResponse(error string) {
	proj := UploadErrorProjection{
		Error: error,
	}

	r.ctx.JSON(500, proj)
}
