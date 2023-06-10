package internal

import (
	"math"

	"gonum.org/v1/gonum/graph/simple"
)

type bridgeDetectionCalculation struct {
	id  float64
	ids map[int64]float64
	low map[int64]float64
}

type bridge struct {
	From int64
	To   int64
}

func DetectBridgesInGraph(graph *simple.UndirectedGraph) []bridge {
	bdc := &bridgeDetectionCalculation{
		id:  0,
		ids: make(map[int64]float64),
		low: make(map[int64]float64),
	}

	visited := make(map[int64]bool)
	bridges := make([]bridge, 0)

	nodes := graph.Nodes()

	for nodes.Next() {
		node := nodes.Node()

		if !visited[node.ID()] {
			bridges = detectBridgesInGraphRecursive(bdc, graph, node.ID(), visited, -1, bridges)
		}
	}

	return bridges
}

func detectBridgesInGraphRecursive(
	bdc *bridgeDetectionCalculation,
	graph *simple.UndirectedGraph,
	at int64,
	visited map[int64]bool,
	parent int64,
	bridges []bridge,
) []bridge {
	visited[at] = true
	bdc.id++
	bdc.low[at] = bdc.id
	bdc.ids[at] = bdc.id

	neighbours := graph.From(at)

	for neighbours.Next() {
		to := neighbours.Node().ID()

		if to == parent {
			continue
		}

		if visited[to] {
			bdc.low[at] = math.Min(bdc.low[at], bdc.ids[to])
		} else {
			bridges = detectBridgesInGraphRecursive(bdc, graph, to, visited, at, bridges)
			bdc.low[at] = math.Min(bdc.low[at], bdc.low[to])

			if bdc.ids[at] < bdc.low[to] {
				if at < to {
					bridges = append(bridges, bridge{
						From: at,
						To:   to,
					})
				} else {
					bridges = append(bridges, bridge{
						From: to,
						To:   at,
					})
				}
			}
		}
	}

	return bridges
}
