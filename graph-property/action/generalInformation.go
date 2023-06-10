package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type GeneralInformationAction interface {
	_interface.Action
}

type generalInformation struct {
	service domain.GeneralInformationService
}

func NewGeneralInformationAction(service domain.GeneralInformationService) GeneralInformationAction {
	return &generalInformation{
		service: service,
	}
}

// Invoke acts as the endpoint for the general information of the graph godoc
//
//	@Router			/graph-property/general-information [get]
//	@Tags			graph-properties
//	@Summary		Get general information
//	@Description	Get general information about the graph (number of nodes, number of edges, network name, creation date)
//	@Produce		json
//	@Success		200	{object}	responder.GeneralInformationProjection
//	@Failure		500
func (a *generalInformation) Invoke(ctx *gin.Context) {
	res := responder.NewGeneralInformationResponder(ctx)

	dto := a.service.Invoke()

	res.CreateSuccessfulResponse(dto)
}
