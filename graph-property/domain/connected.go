package domain

import (
	"graph-analyzer/api/repository"
)

type ConnectedService interface {
	Invoke() ConnectedDTO
}

type connected struct {
	repository repository.GraphRepository
}

func NewConnectedService(repository repository.GraphRepository) ConnectedService {
	return &connected{
		repository: repository,
	}
}

type ConnectedDTO struct {
	Status bool
}

func (s *connected) Invoke() ConnectedDTO {
	result := s.repository.GetConnected()

	return ConnectedDTO{
		Status: result.ComponentCount == 1,
	}
}
