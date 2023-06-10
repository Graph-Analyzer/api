package internal

import (
	"math"

	"gonum.org/v1/gonum/graph/simple"
)

type articulationPointDetectionCalculation struct {
	id                 float64
	ids                map[int64]float64
	low                map[int64]float64
	articulationPoints map[int64]int64
}

func DetectArticulationPointsInGraph(graph *simple.UndirectedGraph) map[int64]int64 {
	apd := &articulationPointDetectionCalculation{
		id:                 0,
		ids:                make(map[int64]float64),
		low:                make(map[int64]float64),
		articulationPoints: make(map[int64]int64),
	}

	visited := make(map[int64]bool)

	nodes := graph.Nodes()

	for nodes.Next() {
		node := nodes.Node()

		if !visited[node.ID()] {
			detectArticulationPointsInGraphRecursive(apd, graph, node.ID(), visited, node.ID(), -1)
		}
	}

	return apd.articulationPoints
}

func detectArticulationPointsInGraphRecursive(
	apd *articulationPointDetectionCalculation,
	graph *simple.UndirectedGraph,
	root int64,
	visited map[int64]bool,
	at int64,
	parent int64,
) {
	visited[at] = true
	apd.id++
	apd.low[at] = apd.id
	apd.ids[at] = apd.id

	children := 0

	neighbours := graph.From(at)

	for neighbours.Next() {
		to := neighbours.Node().ID()

		if to == parent {
			continue
		}

		if visited[to] {
			apd.low[at] = math.Min(apd.low[at], apd.ids[to])
		} else {
			detectArticulationPointsInGraphRecursive(apd, graph, root, visited, to, at)
			apd.low[at] = math.Min(apd.low[at], apd.low[to])

			if apd.ids[at] <= apd.low[to] && parent != -1 {
				apd.articulationPoints[at] = at
			}

			children++
		}
	}

	if parent == -1 && children > 1 {
		apd.articulationPoints[at] = at
	}
}
