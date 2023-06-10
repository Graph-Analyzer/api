package domain

import (
	"graph-analyzer/api/repository"
	"sort"
)

type FullGraphService interface {
	Invoke() FullGraphDTO
}

type fullGraph struct {
	repository repository.GraphRepository
}

func NewFullGraphService(repository repository.GraphRepository) FullGraphService {
	return &fullGraph{
		repository: repository,
	}
}

type FullGraphDTO struct {
	Nodes []fullGraphNodeDTO
	Edges []fullGraphEdgeDTO
}

type fullGraphNodeDTO struct {
	Id    string
	Label string
	Size  int64
}

type fullGraphEdgeDTO struct {
	Id     string
	From   string
	To     string
	Weight float64
}

func (s *fullGraph) Invoke() FullGraphDTO {
	result := s.repository.GetFullGraph()

	nodes := make([]fullGraphNodeDTO, 0)

	nodeKeys := make([]int64, 0, len(result.Nodes))

	for nodeKey := range result.Nodes {
		nodeKeys = append(nodeKeys, nodeKey)
	}

	sort.Slice(
		nodeKeys,
		func(i, j int) bool {
			return nodeKeys[i] < nodeKeys[j]
		},
	)

	for _, nodeKey := range nodeKeys {
		nodes = append(nodes, fullGraphNodeDTO{
			Id:    result.Nodes[nodeKey].RouterId,
			Label: result.Nodes[nodeKey].Label,
			Size:  result.Nodes[nodeKey].Size,
		})
	}

	edges := make([]fullGraphEdgeDTO, 0)

	for _, edge := range result.Edges {
		edges = append(edges, fullGraphEdgeDTO{
			Id:     edge.EdgeId,
			From:   edge.FromRouterId,
			To:     edge.ToRouterId,
			Weight: edge.Weight,
		})
	}

	return FullGraphDTO{
		Nodes: nodes,
		Edges: edges,
	}
}
