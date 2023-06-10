package action

import (
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	_interface "graph-analyzer/api/interface"

	"github.com/gin-gonic/gin"
)

type QueryAction interface {
	_interface.Action
}

type query struct {
	service domain.QueryService
}

func NewQueryAction(service domain.QueryService) QueryAction {
	return &query{
		service: service,
	}
}

type QueryRequestModel struct {
	Type string `form:"type" json:"type" binding:"required" validate:"required,ascii,min=5,max=30" enums:"cut_edges,cut_vertices" example:"cut_edges"`
}

// Invoke acts as the endpoint for queries godoc
//
//	@Router			/graph-property/query [post]
//	@Tags			graph-properties-query
//	@Summary		Get query results
//	@Description	Get results of query on network graph
//	@Accept			json
//	@Produce		json
//	@Param			type	body		action.QueryRequestModel	true	"Type"
//	@Success		200		{object}	responder.QueryProjection
//	@Failure		400
//	@Failure		500
func (a *query) Invoke(ctx *gin.Context) {
	res := responder.NewQueryResponder(ctx)

	var queryRequestModel QueryRequestModel

	err := ctx.BindJSON(&queryRequestModel)
	if err != nil {
		return
	}

	dto := a.service.Invoke(queryRequestModel.Type)

	res.CreateSuccessfulResponse(dto)
}
