package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/simple"
)

func TestDetectArticulationPointsInGraph(t *testing.T) {
	graph := createArticulationPointDetectionTestGraph(false)
	articulationPoints := DetectArticulationPointsInGraph(graph)

	assert.Len(t, articulationPoints, 0)

	graph = createArticulationPointDetectionTestGraph(true)
	articulationPoints = DetectArticulationPointsInGraph(graph)

	assert.Len(t, articulationPoints, 2)
	assert.Equal(t, map[int64]int64{0: 0, 4: 4}, articulationPoints)
}

func createArticulationPointDetectionTestGraph(hasArticulationPoint bool) *simple.UndirectedGraph {
	/**
	           --------------------------
	         /                            \
	       (1)-------------------(4)------(6)
	      /   \                //   \     /
	     /     \   ----------  /     \   /
	    /       \ /      	  /       \ /
	  (2)-------(0)---------(3)       (5)

	Nodes (0) and (4) are articulation points.
	Edges (1)--(4) and (1)--(6) are added if it should not have any articulation points.
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
	node6 := graph.NewNode()
	graph.AddNode(node6)

	// Left triangle
	edge01 := graph.NewEdge(node0, node1)
	graph.SetEdge(edge01)

	edge12 := graph.NewEdge(node1, node2)
	graph.SetEdge(edge12)

	edge20 := graph.NewEdge(node2, node0)
	graph.SetEdge(edge20)

	// Right triangle
	edge45 := graph.NewEdge(node4, node5)
	graph.SetEdge(edge45)

	edge56 := graph.NewEdge(node5, node6)
	graph.SetEdge(edge56)

	edge64 := graph.NewEdge(node6, node4)
	graph.SetEdge(edge64)

	// Triangle connections
	edge03 := graph.NewEdge(node0, node3)
	graph.SetEdge(edge03)

	edge34 := graph.NewEdge(node3, node4)
	graph.SetEdge(edge34)

	edge40 := graph.NewEdge(node4, node0)
	graph.SetEdge(edge40)

	if !hasArticulationPoint {
		edge14 := graph.NewEdge(node1, node4)
		graph.SetEdge(edge14)

		edge16 := graph.NewEdge(node1, node6)
		graph.SetEdge(edge16)
	}

	return graph
}
