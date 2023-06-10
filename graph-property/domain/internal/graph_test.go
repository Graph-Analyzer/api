package internal

import (
	"graph-analyzer/api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/multi"
	"gonum.org/v1/gonum/graph/simple"
)

func TestCreateUnweightedUndirectedSimpleGraphFromFullGraph(t *testing.T) {
	fullGraph := createRepositoryTestGraph()

	createdGraph := CreateUnweightedUndirectedSimpleGraphFromFullGraph(fullGraph)

	assert.IsType(t, &simple.UndirectedGraph{}, createdGraph)
	assert.Equal(t, 3, createdGraph.Nodes().Len())
	assert.Equal(t, 3, createdGraph.Edges().Len())
	assert.True(t, createdGraph.HasEdgeBetween(1, 2))
	assert.True(t, createdGraph.HasEdgeBetween(2, 1))
	assert.True(t, createdGraph.HasEdgeBetween(2, 3))
	assert.True(t, createdGraph.HasEdgeBetween(3, 2))
	assert.True(t, createdGraph.HasEdgeBetween(3, 1))
	assert.True(t, createdGraph.HasEdgeBetween(1, 3))
}

func TestCreateWeightedDirectedMultiGraphFromFullGraph(t *testing.T) {
	fullGraph := createRepositoryTestGraph()

	createdGraph := CreateWeightedDirectedMultiGraphFromFullGraph(fullGraph)

	assert.IsType(t, &multi.WeightedDirectedGraph{}, createdGraph)
	assert.Equal(t, 3, createdGraph.Nodes().Len())
	assert.Equal(t, 5, createdGraph.Edges().Len())
	assert.True(t, createdGraph.HasEdgeFromTo(1, 2))
	assert.True(t, createdGraph.HasEdgeFromTo(2, 1))
	assert.True(t, createdGraph.HasEdgeFromTo(2, 3))
	assert.True(t, createdGraph.HasEdgeFromTo(3, 2))
	assert.True(t, createdGraph.HasEdgeFromTo(3, 1))
	assert.False(t, createdGraph.HasEdgeFromTo(1, 3))
	assert.Equal(t, 2, createdGraph.Lines(3, 1).Len())
	assert.Equal(t, 1.0, createdGraph.WeightedEdge(3, 1).Weight())
}

func createRepositoryTestGraph() repository.FullGraphItem {
	nodes := make(map[int64]repository.FullGraphNodeItem)
	nodes[1] = repository.FullGraphNodeItem{
		Id: 1,
	}
	nodes[2] = repository.FullGraphNodeItem{
		Id: 2,
	}
	nodes[3] = repository.FullGraphNodeItem{
		Id: 3,
	}

	// Bidirectional relations
	edges := make([]repository.FullGraphEdgeItem, 0)
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   1,
		To:     2,
		Weight: 1,
	})
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   2,
		To:     1,
		Weight: 1,
	})
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   2,
		To:     3,
		Weight: 1,
	})
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   3,
		To:     2,
		Weight: 1,
	})
	// Unidirectional
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   3,
		To:     1,
		Weight: 1,
	})
	// Parallel edge for multi graph
	edges = append(edges, repository.FullGraphEdgeItem{
		From:   3,
		To:     1,
		Weight: 2,
	})

	return repository.FullGraphItem{
		Nodes: nodes,
		Edges: edges,
	}
}
