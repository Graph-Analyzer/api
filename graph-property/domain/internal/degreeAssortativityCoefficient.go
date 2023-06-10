package internal

import (
	"math"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/stat"
)

func CalculateDegreeAssortativityCoefficientOfGraph(graph *simple.UndirectedGraph) float64 {
	x := make([]float64, 0)
	y := make([]float64, 0)
	nodes := graph.Nodes()

	for nodes.Next() {
		neighbors := graph.From(nodes.Node().ID())
		nodeDegree := neighbors.Len()

		for neighbors.Next() {
			neighbor := neighbors.Node()
			neighborDegree := graph.From(neighbor.ID()).Len()

			x = append(x, float64(nodeDegree))
			y = append(y, float64(neighborDegree))
		}
	}

	correlation := stat.Correlation(x, y, nil)
	if math.IsNaN(correlation) {
		correlation = 0
	}

	return correlation
}
