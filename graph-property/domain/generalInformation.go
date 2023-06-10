package domain

import (
	"graph-analyzer/api/repository"
	"strings"
	"time"
)

type GeneralInformationService interface {
	Invoke() GeneralInformationDTO
}

type generalInformation struct {
	repository repository.GraphRepository
}

func NewGeneralInformationService(repository repository.GraphRepository) GeneralInformationService {
	return &generalInformation{
		repository: repository,
	}
}

type GeneralInformationDTO struct {
	Name      string
	Nodes     int64
	Edges     int64
	CreatedAt time.Time
}

func (s *generalInformation) Invoke() GeneralInformationDTO {
	result := s.repository.GetGeneralInformation()
	fullGraph := s.repository.GetFullGraph()

	nameSlice := strings.SplitAfter(result.Name, ":")
	name := strings.Join(nameSlice[:len(nameSlice)-1], "")
	name = strings.TrimSuffix(name, ":")

	return GeneralInformationDTO{
		Name:      name,
		Nodes:     int64(len(fullGraph.Nodes)),
		Edges:     int64(len(fullGraph.Edges)),
		CreatedAt: result.CreatedAt,
	}
}
