package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type HasCutVertexService interface {
	Invoke() HasCutVertexDTO
}

type hasCutVertex struct {
	repository repository.GraphRepository
}

func NewHasCutVertexService(repository repository.GraphRepository) HasCutVertexService {
	return &hasCutVertex{
		repository: repository,
	}
}

type HasCutVertexDTO struct {
	Status bool
}

func (s *hasCutVertex) Invoke() HasCutVertexDTO {
	graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	articulationPoints := internal.DetectArticulationPointsInGraph(graph)

	return HasCutVertexDTO{
		Status: len(articulationPoints) > 0,
	}
}
