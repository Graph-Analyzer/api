package action

import (
	mocks "graph-analyzer/api/mocks/repository"
	"graph-analyzer/api/repository"
	"testing"
	"time"
)

func createMockRepository(t *testing.T) repository.GraphRepository {
	mockRepository := mocks.NewGraphRepository(t)

	mockRepository.On("GetAverageClusteringCoefficient").Return(repository.AverageClusteringCoefficientItem{
		Value: 0.77778,
	}).Maybe()

	mockRepository.On("GetConnected").Return(repository.ConnectedItem{
		ComponentCount: 1,
	}).Maybe()

	mockRepository.On("GetDegreeDistribution").Return(repository.DegreeDistributionItem{
		Values: []repository.DegreeDistributionValueItem{
			{
				Degree: 2,
				Amount: 4,
			},
			{
				Degree: 3,
				Amount: 2,
			},
		},
	}).Maybe()

	mockRepository.On("GetDensity").Return(repository.DensityItem{
		Value: 0.46667,
	}).Maybe()

	mockRepository.On("GetLatestGraphName").Return(repository.GraphItem{
		Name: "test:5231bedd-4661-4c91-bc80-3e259b1df941",
	}).Maybe()

	mockRepository.On("GetGeneralInformation").Return(repository.GeneralInformationItem{
		Name:      "test:5231bedd-4661-4c91-bc80-3e259b1df941",
		CreatedAt: time.Date(2022, 11, 30, 13, 59, 59, 0, &time.Location{}),
	}).Maybe()

	/**
	       (3)                  (5)
	      /   \                /   \
	    [1]   [1]            [1]   [1]
	    /       \	         /       \
	  (2)--[1]--(1)--------(4)--[1]--(6)
	*/
	mockRepository.On("GetFullGraph").Return(repository.FullGraphItem{
		Nodes: map[int64]repository.FullGraphNodeItem{
			1: {
				Id:       1,
				RouterId: "1",
				Label:    "XR-1",
				Size:     3,
			},
			2: {
				Id:       2,
				RouterId: "2",
				Label:    "XR-2",
				Size:     2,
			},
			3: {
				Id:       3,
				RouterId: "3",
				Label:    "XR-3",
				Size:     2,
			},
			4: {
				Id:       4,
				RouterId: "4",
				Label:    "XR-4",
				Size:     3,
			},
			5: {
				Id:       5,
				RouterId: "5",
				Label:    "XR-5",
				Size:     2,
			},
			6: {
				Id:       6,
				RouterId: "6",
				Label:    "XR-6",
				Size:     2,
			},
		},
		Edges: []repository.FullGraphEdgeItem{
			{
				EdgeId:       "1",
				From:         1,
				FromRouterId: "1",
				To:           2,
				ToRouterId:   "2",
				Weight:       1,
			},
			{
				EdgeId:       "2",
				From:         2,
				FromRouterId: "2",
				To:           1,
				ToRouterId:   "1",
				Weight:       1,
			},
			{
				EdgeId:       "3",
				From:         2,
				FromRouterId: "2",
				To:           3,
				ToRouterId:   "3",
				Weight:       1,
			},
			{
				EdgeId:       "4",
				From:         3,
				FromRouterId: "3",
				To:           2,
				ToRouterId:   "2",
				Weight:       1,
			},
			{
				EdgeId:       "5",
				From:         3,
				FromRouterId: "3",
				To:           1,
				ToRouterId:   "1",
				Weight:       1,
			},
			{
				EdgeId:       "6",
				From:         1,
				FromRouterId: "1",
				To:           3,
				ToRouterId:   "3",
				Weight:       1,
			},
			{
				EdgeId:       "7",
				From:         4,
				FromRouterId: "4",
				To:           5,
				ToRouterId:   "5",
				Weight:       1,
			},
			{
				EdgeId:       "8",
				From:         5,
				FromRouterId: "5",
				To:           4,
				ToRouterId:   "4",
				Weight:       1,
			},
			{
				EdgeId:       "9",
				From:         5,
				FromRouterId: "5",
				To:           6,
				ToRouterId:   "6",
				Weight:       1,
			},
			{
				EdgeId:       "10",
				From:         6,
				FromRouterId: "6",
				To:           5,
				ToRouterId:   "5",
				Weight:       1,
			},
			{
				EdgeId:       "11",
				From:         6,
				FromRouterId: "6",
				To:           4,
				ToRouterId:   "4",
				Weight:       1,
			},
			{
				EdgeId:       "12",
				From:         4,
				FromRouterId: "4",
				To:           6,
				ToRouterId:   "6",
				Weight:       1,
			},
			{
				EdgeId:       "13",
				From:         1,
				FromRouterId: "1",
				To:           4,
				ToRouterId:   "4",
				Weight:       1,
			},
			{
				EdgeId:       "14",
				From:         4,
				FromRouterId: "4",
				To:           1,
				ToRouterId:   "1",
				Weight:       1,
			},
		},
	}).Maybe()

	return mockRepository
}
