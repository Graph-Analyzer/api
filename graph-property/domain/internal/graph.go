package internal

import (
	"graph-analyzer/api/repository"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
	"gonum.org/v1/gonum/graph/simple"
)

func createEmptyUnweightedUndirectedSimpleGraph() *simple.UndirectedGraph {
	return simple.NewUndirectedGraph()
}

func CreateUnweightedUndirectedSimpleGraphFromFullGraph(fullGraph repository.FullGraphItem) *simple.UndirectedGraph {
	unweightedUndirectedSimpleGraph := createEmptyUnweightedUndirectedSimpleGraph()

	for _, node := range fullGraph.Nodes {
		unweightedUndirectedSimpleGraph.AddNode(simple.Node(node.Id))
	}

	for _, edge := range fullGraph.Edges {
		edge := unweightedUndirectedSimpleGraph.NewEdge(
			unweightedUndirectedSimpleGraph.Node(edge.From),
			unweightedUndirectedSimpleGraph.Node(edge.To),
		)

		unweightedUndirectedSimpleGraph.SetEdge(edge)
	}

	return unweightedUndirectedSimpleGraph
}

func createEmptyWeightedDirectedMultiGraph() *multi.WeightedDirectedGraph {
	weightedDirectedMultiGraph := multi.NewWeightedDirectedGraph()

	// Custom EdgeWeight function to only consider the lowest line weight per edge
	weightedDirectedMultiGraph.EdgeWeightFunc = func(lines graph.WeightedLines) float64 {
		if lines.Len() == 0 {
			return 0
		}
		var weight float64
		first := true

		for lines.Next() {
			if lineWeight := lines.WeightedLine().Weight(); first || lineWeight < weight {
				weight = lineWeight
			}

			first = false
		}

		lines.Reset()

		return weight
	}

	return weightedDirectedMultiGraph
}

func CreateWeightedDirectedMultiGraphFromFullGraph(fullGraph repository.FullGraphItem) *multi.WeightedDirectedGraph {
	weightedDirectedMultiGraph := createEmptyWeightedDirectedMultiGraph()

	for _, node := range fullGraph.Nodes {
		weightedDirectedMultiGraph.AddNode(multi.Node(node.Id))
	}

	for _, edge := range fullGraph.Edges {
		weightedLine := weightedDirectedMultiGraph.NewWeightedLine(
			weightedDirectedMultiGraph.Node(edge.From),
			weightedDirectedMultiGraph.Node(edge.To), edge.Weight,
		)

		weightedDirectedMultiGraph.SetWeightedLine(weightedLine)
	}

	return weightedDirectedMultiGraph
}
