package problem1

import (
	"testing"

	"github.com/manofthelionarmy/goLearnAlgos/graphs"
	"github.com/stretchr/testify/require"
)

func TestRoute(t *testing.T) {
	g := graphs.NewGraph()
	g.AddEdge(1, 5)
	g.AddEdge(5, 2)
	g.AddEdge(2, 6)
	g.AddEdge(1, 4)

	verticies := g.Vertices()
	source := verticies[1]
	end := verticies[6]

	// I assumed that we arleady have a bfs tree produced after the bfs search/traversal.
	// we must create a path with a given source vertex, and check if the given end vertex is on the path
	// So, let's perform the bfs. Part of this algorithm is to 'clean the slate' and a new tree is produced give a new source vertex
	require.Equal(t, true, IsThereARoute(g, source, end))

	require.Equal(t, true, IsThereARoute(g, verticies[4], verticies[6]))
	// Just because the vertex is in the adjaceny list doesn't mean it's in the path.

	// A path is comprised of edges. Edges are developed by perpetually visiting verticies in each adjancency list.
	// A breadth first search tree is produced with this as the root, or source.
	require.Equal(t, true, SolutionIsEnRoute(g, source, end))
	require.Equal(t, true, SolutionIsEnRoute(g, verticies[4], verticies[6]))
}
