package domain

import (
	"graph-analyzer/api/repository"
)

type AverageClusteringCoefficientService interface {
	Invoke() AverageClusteringCoefficientDTO
}

type averageClusteringCoefficient struct {
	repository repository.GraphRepository
}

func NewAverageClusteringCoefficientService(repository repository.GraphRepository) AverageClusteringCoefficientService {
	return &averageClusteringCoefficient{
		repository: repository,
	}
}

type AverageClusteringCoefficientDTO struct {
	Value float64
}

func (s *averageClusteringCoefficient) Invoke() AverageClusteringCoefficientDTO {
	result := s.repository.GetAverageClusteringCoefficient()

	return AverageClusteringCoefficientDTO{
		Value: result.Value,
	}
}
