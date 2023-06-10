package domain

import (
	"graph-analyzer/api/graph-property/domain/internal"
	"graph-analyzer/api/repository"
)

type DegreeCorrelationService interface {
	Invoke() DegreeCorrelationDTO
}

type degreeCorrelation struct {
	repository repository.GraphRepository
}

func NewDegreeCorrelationService(repository repository.GraphRepository) DegreeCorrelationService {
	return &degreeCorrelation{
		repository: repository,
	}
}

type DegreeCorrelationDTO struct {
	Values []degreeCorrelationValueDTO
}

type degreeCorrelationValueDTO struct {
	Degree  int64
	Average float64
}

func (s *degreeCorrelation) Invoke() DegreeCorrelationDTO {
	graph := internal.CreateUnweightedUndirectedSimpleGraphFromFullGraph(
		s.repository.GetFullGraph(),
	)

	degreeCorrelations := internal.CalculateDegreeCorrelationOfGraph(graph)

	values := make([]degreeCorrelationValueDTO, 0)

	for _, degreeCorrelation := range degreeCorrelations {
		values = append(values, degreeCorrelationValueDTO{
			Degree:  degreeCorrelation.Degree,
			Average: degreeCorrelation.Average,
		})
	}

	return DegreeCorrelationDTO{
		Values: values,
	}
}
