package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/simple"
)

func TestCalculateDegreeCorrelationOfGraph(t *testing.T) {
	graph := createDegreeCorrelationTestGraph()
	degreeCorrelation := CalculateDegreeCorrelationOfGraph(graph)

	assert.Len(t, degreeCorrelation, 2)
	assert.Equal(t, int64(2), degreeCorrelation[0].Degree)
	assert.InDelta(t, 2.5, degreeCorrelation[0].Average, 0)
	assert.Equal(t, int64(3), degreeCorrelation[1].Degree)
	assert.InDelta(t, 2.3333, degreeCorrelation[1].Average, 0.0001)
}

func createDegreeCorrelationTestGraph() *simple.UndirectedGraph {
	/**
	       (1)                   (4)
	      /   \                 /   \
	     /     \               /     \
	    /       \	      	  /       \
	  (2)-------(0)---------(3)-------(5)
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
