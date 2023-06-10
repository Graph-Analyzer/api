package responder

import (
	"graph-analyzer/api/graph-property/domain"

	"github.com/gin-gonic/gin"
)

type FullGraphResponder interface {
	CreateSuccessfulResponse(dto domain.FullGraphDTO)
}

type fullGraph struct {
	ctx *gin.Context
}

func NewFullGraphResponder(ctx *gin.Context) FullGraphResponder {
	return &fullGraph{
		ctx: ctx,
	}
}

type FullGraphProjection struct {
	Nodes []FullGraphNodeProjection `json:"nodes" validate:"required"`
	Edges []FullGraphEdgeProjection `json:"edges" validate:"required"`
}

type FullGraphNodeProjection struct {
	Id    string `json:"id" validate:"required" example:"6"`
	Label string `json:"label" validate:"required" example:"XR-6"`
	Size  int64  `json:"size" validate:"required" example:"5"`
}

type FullGraphEdgeProjection struct {
	Id     string  `json:"id" validate:"required" example:"12"`
	From   string  `json:"from" validate:"required" example:"6"`
	To     string  `json:"to" validate:"required" example:"5"`
	Weight float64 `json:"weight" validate:"required" example:"1.0"`
}

func (r *fullGraph) CreateSuccessfulResponse(dto domain.FullGraphDTO) {
	nodes := make([]FullGraphNodeProjection, 0)

	for _, node := range dto.Nodes {
		nodes = append(nodes, FullGraphNodeProjection{
			Id:    node.Id,
			Label: node.Label,
			Size:  node.Size,
		})
	}

	edges := make([]FullGraphEdgeProjection, 0)

	for _, edge := range dto.Edges {
		edges = append(edges, FullGraphEdgeProjection{
			Id:     edge.Id,
			From:   edge.From,
			To:     edge.To,
			Weight: edge.Weight,
		})
	}

	proj := FullGraphProjection{
		Nodes: nodes,
		Edges: edges,
	}

	r.ctx.JSON(200, proj)
}
