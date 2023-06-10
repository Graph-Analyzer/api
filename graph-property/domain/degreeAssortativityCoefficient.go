package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type DegreeAssortativityCoefficientService interface {
	Invoke() DegreeAssortativityCoefficientDTO
}

type degreeAssortativityCoefficient struct {
	repository repository.GraphRepository
}

func NewDegreeAssortativityCoefficientService(repository repository.GraphRepository) DegreeAssortativityCoefficientService {
	return &degreeAssortativityCoefficient{
		repository: repository,
	}
}

type DegreeAssortativityCoefficientDTO struct {
	Value float64
}

func (s *degreeAssortativityCoefficient) Invoke() DegreeAssortativityCoefficientDTO {
	graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	value := internal.CalculateDegreeAssortativityCoefficientOfGraph(graph)

	return DegreeAssortativityCoefficientDTO{
		Value: value,
	}
}
