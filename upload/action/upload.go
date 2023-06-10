package action

import (
	_interface "graph-analyzer/api/interface"
	"graph-analyzer/api/upload/domain"
	"graph-analyzer/api/upload/responder"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadAction interface {
	_interface.Action
}

type upload struct {
	service domain.UploadService
}

func NewUploadAction(service domain.UploadService) UploadAction {
	return &upload{
		service: service,
	}
}

type UploadRequestModel struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// Invoke acts as the endpoint for the upload endpoint godoc
//
//	@Router			/upload [post]
//	@Tags			upload
//	@Summary		Upload gexf
//	@Description	Upload graph-file (gexf)
//	@Param			file	formData	file	true	"GEXF file"
//	@Produce		json
//	@Success		200	{object}	responder.UploadProjection
//	@Failure		400	{object}	responder.UploadErrorProjection
//	@Failure		500	{object}	responder.UploadErrorProjection
func (a *upload) Invoke(ctx *gin.Context) {
	// Limit Request Body to 8MiB
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 8<<20)

	res := responder.NewUploadResponder(ctx)

	var uploadRequestModel UploadRequestModel

	if err := ctx.ShouldBind(&uploadRequestModel); err != nil {
		res.CreateBadRequestErrorResponse(err.Error())
		return
	}

	if !strings.HasSuffix(uploadRequestModel.File.Filename, ".gexf") {
		res.CreateBadRequestErrorResponse("File extension not supported")
		return
	}

	dto, err := a.service.Invoke(uploadRequestModel.File)
	if err != nil {
		// TODO: Possible leaking sensitive data
		res.CreateInternalErrorResponse(err.Error())
		return
	}

	res.CreateSuccessfulResponse(dto)
}
