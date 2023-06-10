package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type HasCutEdgeService interface {
	Invoke() HasCutEdgeDTO
}

type hasCutEdge struct {
	repository repository.GraphRepository
}

func NewHasCutEdgeService(repository repository.GraphRepository) HasCutEdgeService {
	return &hasCutEdge{
		repository: repository,
	}
}

type HasCutEdgeDTO struct {
	Status bool
}

func (s *hasCutEdge) Invoke() HasCutEdgeDTO {
	graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	bridges := internal.DetectBridgesInGraph(graph)

	return HasCutEdgeDTO{
		Status: len(bridges) > 0,
	}
}
