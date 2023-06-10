package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/multi"
)

func TestCalculateDiameterOfGraph(t *testing.T) {
	graph := createDiameterTestGraph(false)
	diameter, hops := CalculateDiameterOfGraph(graph)

	assert.Equal(t, 13.5, diameter)
	assert.Equal(t, int64(4), hops)
}

func TestCalculateDiameterOfGraphWithNegativeWeight(t *testing.T) {
	graph := createDiameterTestGraph(true)
	diameter, hops := CalculateDiameterOfGraph(graph)

	assert.Equal(t, 0.0, diameter)
	assert.Equal(t, int64(0), hops)
}

func createDiameterTestGraph(negativeWeight bool) *multi.WeightedDirectedGraph {
	/**
	       (1)               -->(4)
	      /   \            /   /   \
	    [1]   [1]      [1.5] [10]  [1]
	    /       \	       \ /       \
	  (0)<-[1]--(2)--[10]--(3)--[1]--(5)

	All bidirectional except (0)<-(2) and (3)->(4), those are unidirectional
	Additional parallel line between (3) and (4) that is unidirectional
	Longest shortest path is between (0) and (4)
	*/

	graph := createEmptyWeightedDirectedMultiGraph()

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
	line01 := graph.NewWeightedLine(node0, node1, 1)
	graph.SetWeightedLine(line01)

	line10 := graph.NewWeightedLine(node1, node0, 1)
	graph.SetWeightedLine(line10)

	line12 := graph.NewWeightedLine(node1, node2, 1)
	graph.SetWeightedLine(line12)

	line21 := graph.NewWeightedLine(node2, node1, 1)
	graph.SetWeightedLine(line21)

	line20 := graph.NewWeightedLine(node2, node0, 1) // Unidirectional
	graph.SetWeightedLine(line20)

	line34Weight := 1.5

	if negativeWeight {
		line34Weight = -1
	}

	// Right triangle
	line34Parallel := graph.NewWeightedLine(node3, node4, line34Weight) // Parallel unidirectional line
	graph.SetWeightedLine(line34Parallel)                               // Set parallel edge with lesser weight first to test lack of override

	line34 := graph.NewWeightedLine(node3, node4, 10)
	graph.SetWeightedLine(line34)

	line43 := graph.NewWeightedLine(node4, node3, 10)
	graph.SetWeightedLine(line43)

	line45 := graph.NewWeightedLine(node4, node5, 1)
	graph.SetWeightedLine(line45)

	line54 := graph.NewWeightedLine(node5, node4, 1)
	graph.SetWeightedLine(line54)

	line53 := graph.NewWeightedLine(node5, node3, 1)
	graph.SetWeightedLine(line53)

	line35 := graph.NewWeightedLine(node3, node5, 1)
	graph.SetWeightedLine(line35)

	// Triangle connections
	line23 := graph.NewWeightedLine(node2, node3, 10)
	graph.SetWeightedLine(line23)

	line32 := graph.NewWeightedLine(node3, node2, 10)
	graph.SetWeightedLine(line32)

	return graph
}
