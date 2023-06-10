package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type DiameterService interface {
	Invoke() DiameterDTO
}

type diameter struct {
	repository repository.GraphRepository
}

func NewDiameterService(repository repository.GraphRepository) DiameterService {
	return &diameter{
		repository: repository,
	}
}

type DiameterDTO struct {
	Diameter float64
	Hops     int64
}

func (s *diameter) Invoke() DiameterDTO {
	graph := internal.CreateWeightedDirectedMultiGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	diameter, hops := internal.CalculateDiameterOfGraph(graph)

	return DiameterDTO{
		Diameter: diameter,
		Hops:     hops,
	}
}
