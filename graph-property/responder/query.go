package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type QueryResponder interface {
	CreateSuccessfulResponse(domain.QueryDTO)
}

type query struct {
	ctx *gin.Context
}

func NewQueryResponder(ctx *gin.Context) QueryResponder {
	return &query{
		ctx: ctx,
	}
}

type QueryProjection struct {
	Nodes []QueryNodeProjection `json:"nodes" validate:"required"`
	Edges []QueryEdgeProjection `json:"edges" validate:"required"`
}

type QueryNodeProjection struct {
	Id    string `json:"id" validate:"required" example:"6"`
	Label string `json:"label" validate:"required" example:"XR-6"`
}

type QueryEdgeProjection struct {
	From      string `json:"from" validate:"required" example:"6"`
	FromLabel string `json:"fromLabel" validate:"required" example:"XR-6"`
	To        string `json:"to" validate:"required" example:"5"`
	ToLabel   string `json:"toLabel" validate:"required" example:"XR-5"`
}

func (r *query) CreateSuccessfulResponse(dto domain.QueryDTO) {
	nodes := make([]QueryNodeProjection, 0)

	for _, node := range dto.Nodes {
		nodes = append(nodes, QueryNodeProjection{
			Id:    node.Id,
			Label: node.Label,
		})
	}

	edges := make([]QueryEdgeProjection, 0)

	for _, edge := range dto.Edges {
		edges = append(edges, QueryEdgeProjection{
			From:      edge.From,
			FromLabel: edge.FromLabel,
			To:        edge.To,
			ToLabel:   edge.ToLabel,
		})
	}

	proj := QueryProjection{
		Nodes: nodes,
		Edges: edges,
	}

	r.ctx.JSON(200, proj)
}
