package internal

import (
	"sort"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/stat"
)

type degreeCorrelationValue struct {
	Degree  int64
	Average float64
}

func CalculateDegreeCorrelationOfGraph(graph *simple.UndirectedGraph) []degreeCorrelationValue {
	var nodeNeighborDegrees = make(map[int64][]float64)
	nodes := graph.Nodes()

	for nodes.Next() {
		neighbors := graph.From(nodes.Node().ID()) // Get all neighbors of node X
		nodeDegree := int64(neighbors.Len())       // Get degree of node X

		for neighbors.Next() { // For neighbor of node X
			neighbor := neighbors.Node()                               // Get id of neighbor of node X
			neighborDegree := float64(graph.From(neighbor.ID()).Len()) // Get degree of neighbor of node X

			nodeNeighborDegrees[nodeDegree] = append(nodeNeighborDegrees[nodeDegree], neighborDegree) // Add to map
		}
	}

	nodeDegrees := make([]int64, 0, len(nodeNeighborDegrees))

	for nodeDegree := range nodeNeighborDegrees {
		nodeDegrees = append(nodeDegrees, nodeDegree)
	}

	sort.Slice(
		nodeDegrees,
		func(i, j int) bool {
			return nodeDegrees[i] < nodeDegrees[j]
		},
	)

	values := make([]degreeCorrelationValue, 0, len(nodeDegrees))

	for _, nodeDegree := range nodeDegrees {
		values = append(values, degreeCorrelationValue{
			Degree:  nodeDegree,
			Average: stat.Mean(nodeNeighborDegrees[nodeDegree], nil),
		})
	}

	return values
}
