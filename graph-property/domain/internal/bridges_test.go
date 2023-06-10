package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/simple"
)

func TestDetectBridgesInGraph(t *testing.T) {
	graph := createBridgeDetectionTestGraph(false)
	bridges := DetectBridgesInGraph(graph)

	assert.Len(t, bridges, 0)

	graph = createBridgeDetectionTestGraph(true)
	bridges = DetectBridgesInGraph(graph)

	assert.Len(t, bridges, 1)
	assert.Equal(t, bridge{From: 0, To: 3}, bridges[0])
}

func createBridgeDetectionTestGraph(hasBridge bool) *simple.UndirectedGraph {
	/**
	       (1)-------------------(4)
	      /   \                 /   \
	     /     \               /     \
	    /       \	      	  /       \
	  (2)-------(0)---------(3)-------(5)

	Edge (0)--(3) is the bridge.
	The edge (1)--(4) is added if it should not have a bridge.
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

	// Triangle connections
	edge03 := graph.NewEdge(node0, node3)
	graph.SetEdge(edge03)

	if !hasBridge {
		edge14 := graph.NewEdge(node1, node4)
		graph.SetEdge(edge14)

		edge41 := graph.NewEdge(node4, node1)
		graph.SetEdge(edge41)
	}

	return graph
}
