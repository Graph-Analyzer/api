package internal

import (
	"log"

	"gonum.org/v1/gonum/graph/multi"
	"gonum.org/v1/gonum/graph/path"
)

func CalculateDiameterOfGraph(graph *multi.WeightedDirectedGraph) (float64, int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred during diameter calculation:", err)
		}
	}()

	shortestPaths := path.DijkstraAllPaths(graph)

	nodesFrom := graph.Nodes()
	nodesTo := graph.Nodes()
	distance := float64(0)
	hops := int64(0)

	for nodesFrom.Next() {
		nodeFrom := nodesFrom.Node()

		for nodesTo.Next() {
			nodeTo := nodesTo.Node()
			shortestPaths, weight := shortestPaths.AllBetween(nodeFrom.ID(), nodeTo.ID())

			if weight >= distance {
				distance = weight

				for _, shortestPath := range shortestPaths {
					if shortestPathHops := int64(len(shortestPath)) - 1; shortestPathHops > hops {
						hops = shortestPathHops
					}
				}

			}
		}

		nodesTo.Reset()
	}

	return distance, hops
}
