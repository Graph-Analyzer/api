package responder

import (
	"graph-analyzer/api/graph-property/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type GeneralInformationResponder interface {
	CreateSuccessfulResponse(dto domain.GeneralInformationDTO)
}

type generalInformation struct {
	ctx *gin.Context
}

func NewGeneralInformationResponder(ctx *gin.Context) GeneralInformationResponder {
	return &generalInformation{
		ctx: ctx,
	}
}

type GeneralInformationProjection struct {
	Name      string `json:"name" validate:"required" example:"SA-Network"`
	Nodes     int64  `json:"nodes" validate:"required" example:"8"`
	Edges     int64  `json:"edges" validate:"required" example:"26"`
	CreatedAt string `json:"created_at" validate:"required" example:"2022-11-10T12:10:07Z"`
}

func (r *generalInformation) CreateSuccessfulResponse(dto domain.GeneralInformationDTO) {
	proj := GeneralInformationProjection{
		Name:      dto.Name,
		Nodes:     dto.Nodes,
		Edges:     dto.Edges,
		CreatedAt: dto.CreatedAt.Format(time.RFC3339),
	}

	r.ctx.JSON(200, proj)
}
