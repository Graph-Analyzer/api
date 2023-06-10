package domain

import (
	"graph-analyzer/api/repository"
)

type DensityService interface {
	Invoke() DensityDTO
}

type density struct {
	repository repository.GraphRepository
}

func NewDensityService(repository repository.GraphRepository) DensityService {
	return &density{
		repository: repository,
	}
}

type DensityDTO struct {
	Value float64
}

func (s *density) Invoke() DensityDTO {
	result := s.repository.GetDensity()

	return DensityDTO{
		Value: result.Value,
	}
}
