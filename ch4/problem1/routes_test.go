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

	graphs.BreadthFirstSearch(g, source)
	require.Equal(t, true, IsThereARoute(source, end))

	require.Equal(t, false, IsThereARoute(verticies[4], verticies[6]))
}
