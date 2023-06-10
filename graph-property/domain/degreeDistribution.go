package domain

import (
	"graph-analyzer/api/repository"
)

type DegreeDistributionService interface {
	Invoke() DegreeDistributionDTO
}

type degreeDistribution struct {
	repository repository.GraphRepository
}

func NewDegreeDistributionService(repository repository.GraphRepository) DegreeDistributionService {
	return &degreeDistribution{
		repository: repository,
	}
}

type DegreeDistributionDTO struct {
	Values []degreeDistributionValueDTO
}

type degreeDistributionValueDTO struct {
	Degree int64
	Amount int64
}

func (s *degreeDistribution) Invoke() DegreeDistributionDTO {
	result := s.repository.GetDegreeDistribution()

	values := make([]degreeDistributionValueDTO, 0)

	for _, value := range result.Values {
		values = append(values, degreeDistributionValueDTO{
			Degree: value.Degree,
			Amount: value.Amount,
		})
	}

	return DegreeDistributionDTO{
		Values: values,
	}
}
