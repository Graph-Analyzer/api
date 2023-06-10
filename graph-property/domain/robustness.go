package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type RobustnessService interface {
	Invoke() RobustnessDTO
}

type robustness struct {
	repository repository.GraphRepository
}

func NewRobustnessService(repository repository.GraphRepository) RobustnessService {
	return &robustness{
		repository: repository,
	}
}

type RobustnessDTO struct {
	Nodes      int64
	Percentage float64
}

func (s *robustness) Invoke() RobustnessDTO {
	graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	nodes, percentage := internal.CalculateRobustnessOfGraph(graph)

	return RobustnessDTO{
		Nodes:      nodes,
		Percentage: percentage,
	}
}
