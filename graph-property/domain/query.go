package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type QueryService interface {
	Invoke(queryType string) QueryDTO
}

type query struct {
	repository repository.GraphRepository
}

func NewQueryService(repository repository.GraphRepository) QueryService {
	return &query{
		repository: repository,
	}
}

const cutEdgesQuery string = "cut_edges"
const cutVerticesQuery string = "cut_vertices"

type QueryDTO struct {
	Nodes []queryNodeDTO
	Edges []queryEdgeDTO
}

type queryNodeDTO struct {
	Id    string
	Label string
}

type queryEdgeDTO struct {
	From      string
	FromLabel string
	To        string
	ToLabel   string
}

func (s *query) Invoke(queryType string) QueryDTO {
	nodes := make([]queryNodeDTO, 0)
	edges := make([]queryEdgeDTO, 0)

	switch queryType {
	case cutEdgesQuery:
		fullGraph := s.repository.GetFullGraph()
		graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(fullGraph)
		bridges := internal.DetectBridgesInGraph(graph)

		for _, bridge := range bridges {
			edges = append(edges, queryEdgeDTO{
				From:      fullGraph.Nodes[bridge.From].RouterId,
				FromLabel: fullGraph.Nodes[bridge.From].Label,
				To:        fullGraph.Nodes[bridge.To].RouterId,
				ToLabel:   fullGraph.Nodes[bridge.To].Label,
			})
		}
	case cutVerticesQuery:
		fullGraph := s.repository.GetFullGraph()
		graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(fullGraph)
		articulationPoints := internal.DetectArticulationPointsInGraph(graph)

		for _, articulationPointId := range articulationPoints {
			nodes = append(nodes, queryNodeDTO{
				Id:    fullGraph.Nodes[articulationPointId].RouterId,
				Label: fullGraph.Nodes[articulationPointId].Label,
			})
		}
	}

	return QueryDTO{
		Nodes: nodes,
		Edges: edges,
	}
}
