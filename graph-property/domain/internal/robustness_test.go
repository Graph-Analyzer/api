package internal

import (
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/simple"

	"testing"
)

func TestCalculateRobustnessOfGraph(t *testing.T) {
	graph := createRobustnessTestGraph(true)
	nodes, percentage := CalculateRobustnessOfGraph(graph)
	assert.Equal(t, int64(0), nodes)
	assert.Equal(t, float64(0), percentage)

	graph = createRobustnessTestGraph(false)
	nodes, percentage = CalculateRobustnessOfGraph(graph)
	assert.Equal(t, int64(1), nodes)
	assert.Equal(t, float64(25), percentage)
}

func createRobustnessTestGraph(hasArticulationPoint bool) *simple.UndirectedGraph {
	/**
	            (3)
	          /  |  \
	        /    |    \
	      /      |      \
	    /        |        \
	  (0)-------(1)-------(2)

	Node (1) is the articulation point (cut-node).
	The edges (0)--(3) and (2)--(3) are added if it should not have an articulation point (cut-node).
	*/

	graph := createEmptyUnweightedUndirectedSimpleGraph()

	node0 := graph.NewNode()
	graph.AddNode(node0)
	node1 := graph.NewNode()
	graph.AddNode(node1)
	node2 := graph.NewNode()
	graph.AddNode(node2)
	node3 := graph.NewNode()
	graph.AddNode(node3)

	edge01 := graph.NewEdge(node0, node1)
	graph.SetEdge(edge01)

	edge02 := graph.NewEdge(node1, node2)
	graph.SetEdge(edge02)

	edge03 := graph.NewEdge(node1, node3)
	graph.SetEdge(edge03)

	if !hasArticulationPoint {
		edge04 := graph.NewEdge(node0, node3)
		graph.SetEdge(edge04)

		edge05 := graph.NewEdge(node2, node3)
		graph.SetEdge(edge05)
	}

	return graph
}
