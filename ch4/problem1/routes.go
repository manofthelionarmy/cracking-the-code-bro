package problem1

import (
	"github.com/manofthelionarmy/goLearnAlgos/graphs"
)

func IsThereARoute(sourceVertex *graphs.Vertex, endVertex *graphs.Vertex) bool {

	trailingVertex := endVertex

	for trailingVertex != nil {
		if sourceVertex == trailingVertex {
			return true
		}
		trailingVertex = trailingVertex.Parent
	}
	return false
}
