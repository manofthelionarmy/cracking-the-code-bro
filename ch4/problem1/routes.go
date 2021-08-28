package problem1

import (
	"github.com/manofthelionarmy/goLearnAlgos/graphs"
)

func IsThereARoute(g graphs.Graph, source *graphs.Vertex, endVertex *graphs.Vertex) bool {
	graphs.BreadthFirstSearch(g, source)

	trailingVertex := endVertex

	for trailingVertex != nil {
		if source == trailingVertex {
			return true
		}
		trailingVertex = trailingVertex.Parent
	}
	return false
}
