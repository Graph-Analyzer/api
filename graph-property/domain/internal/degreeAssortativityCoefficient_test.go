package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/simple"
)

func TestCalculateDegreeAssortativityCoefficientOfGraph(t *testing.T) {
	graph := createDegreeAssortativityCoefficientTestGraph(false)
	degreeAssortativityCoefficient := CalculateDegreeAssortativityCoefficientOfGraph(graph)

	assert.Equal(t, float64(0), degreeAssortativityCoefficient)

	graph = createDegreeAssortativityCoefficientTestGraph(true)
	degreeAssortativityCoefficient = CalculateDegreeAssortativityCoefficientOfGraph(graph)

	assert.InDelta(t, -0.16666, degreeAssortativityCoefficient, 0.00001)
}

func createDegreeAssortativityCoefficientTestGraph(isCorrelated bool) *simple.UndirectedGraph {
	/**
	       (1)                   (4)
	      /   \                 /   \
	     /     \               /     \
	    /       \	      	  /       \
	  (2)-------(0)---------(3)-------(5)
	*/

	graph := createEmptyUnweightedUndirectedSimpleGraph()

	if !isCorrelated {
		node0 := graph.NewNode()
		graph.AddNode(node0)
		node1 := graph.NewNode()
		graph.AddNode(node1)

		edge01 := graph.NewEdge(node0, node1)
		graph.SetEdge(edge01)

		return graph
	}

	node0 := graph.NewNode()
	graph.AddNode(node0)
	node1 := graph.NewNode()
	graph.AddNode(node1)
	node2 := graph.NewNode()
	graph.AddNode(node2)
	node3 := graph.NewNode()
	graph.AddNode(node3)
	node4 := graph.NewNode()
	graph.AddNode(node4)
	node5 := graph.NewNode()
	graph.AddNode(node5)

	// Left triangle
	edge01 := graph.NewEdge(node0, node1)
	graph.SetEdge(edge01)

	edge12 := graph.NewEdge(node1, node2)
	graph.SetEdge(edge12)

	edge20 := graph.NewEdge(node2, node0)
	graph.SetEdge(edge20)

	// Right triangle
	edge34 := graph.NewEdge(node3, node4)
	graph.SetEdge(edge34)

	edge45 := graph.NewEdge(node4, node5)
	graph.SetEdge(edge45)

	edge53 := graph.NewEdge(node5, node3)
	graph.SetEdge(edge53)

	// Triangle connection
	edge03 := graph.NewEdge(node0, node3)
	graph.SetEdge(edge03)

	return graph
}
