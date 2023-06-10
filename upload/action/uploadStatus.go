package action

import (
	_interface "graph-analyzer/api/interface"
	"graph-analyzer/api/upload/domain"
	"graph-analyzer/api/upload/responder"

	"github.com/gin-gonic/gin"
)

type UploadStatusAction interface {
	_interface.Action
}

type uploadStatus struct {
	service domain.UploadStatusService
}

func NewUploadStatusAction(service domain.UploadStatusService) UploadStatusAction {
	return &uploadStatus{
		service: service,
	}
}

// Invoke acts as the endpoint for the status of data-collector endpoint godoc
//
//	@Router			/upload-status [get]
//	@Tags			upload
//	@Summary		Get data-collector status
//	@Description	Get connected status of the data-collector gexf endpoint
//	@Produce		json
//	@Success		200	{object}	responder.UploadStatusProjection
func (a *uploadStatus) Invoke(ctx *gin.Context) {
	res := responder.NewUploadStatusResponder(ctx)

	dto, _ := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
