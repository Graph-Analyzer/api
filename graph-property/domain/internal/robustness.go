package internal

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func CalculateRobustnessOfGraph(fullGraph *simple.UndirectedGraph) (int64, float64) {
	numberOfNodes := int64(fullGraph.Nodes().Len())
	removalMinimum := findNodeRemovalMinimum(fullGraph)

	removalPercentage := float64(removalMinimum) / float64(numberOfNodes) * 100

	return removalMinimum, removalPercentage
}

// Minimum is the amount of nodes to remove until the next removal
// leads to a disconnected graph
func findNodeRemovalMinimum(currentGraph *simple.UndirectedGraph) int64 {
	minimum := findNodeRemovalMinimumRecursive(
		currentGraph,
		int64(currentGraph.Nodes().Len()),
		0,
	)

	if 0 == minimum {
		return minimum
	}

	return minimum - 1
}

// Finds the amount of nodes to remove until the graph is disconnected
func findNodeRemovalMinimumRecursive(
	currentGraph *simple.UndirectedGraph,
	removalMinimum int64,
	currentRemovalAmount int64,
) int64 {
	// Check if graph is split, return the minimum of node removals if this is the case
	if 1 < len(topo.ConnectedComponents(currentGraph)) {
		return currentRemovalAmount
	}

	// Go over all nodes of the currently highest degree
	for _, node := range getHighestDegreeNodes(currentGraph) {
		// Terminate execution if no other minimums are possible
		if removalMinimum <= currentRemovalAmount+1 {
			break
		}

		copyGraph := simple.NewUndirectedGraph()
		graph.Copy(copyGraph, currentGraph)

		copyGraph.RemoveNode(node.ID())

		localRemovalMinimum := findNodeRemovalMinimumRecursive(
			copyGraph,
			removalMinimum,
			currentRemovalAmount+1,
		)

		if localRemovalMinimum < removalMinimum {
			removalMinimum = localRemovalMinimum
		}
	}

	return removalMinimum
}

func getHighestDegreeNodes(currentGraph graph.Undirected) []graph.Node {
	var highestDegreeNodes []graph.Node
	nodes := currentGraph.Nodes()
	highestDegree := int64(0)

	for nodes.Next() {
		neighbors := currentGraph.From(nodes.Node().ID())
		nodeDegree := int64(neighbors.Len())

		if nodeDegree > highestDegree {
			highestDegree = nodeDegree
			highestDegreeNodes = []graph.Node{nodes.Node()}
		} else if nodeDegree == highestDegree {
			highestDegreeNodes = append(highestDegreeNodes, nodes.Node())
		}
	}

	return highestDegreeNodes
}
