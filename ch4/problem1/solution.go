package problem1

import (
	"math"

	"github.com/manofthelionarmy/goLearnAlgos/graphs"
	"github.com/manofthelionarmy/goLearnAlgos/graphs/colors"
	"github.com/manofthelionarmy/goLearnAlgos/queues"
)

func SolutionIsEnRoute(g graphs.Graph, source *graphs.Vertex, end *graphs.Vertex) bool {
	for _, v := range g.Vertices() {
		if v != source {
			v.Color = colors.White
			v.Parent = nil
			v.Distance = math.MaxInt64
		}
	}

	source.Color = colors.Gray
	source.Distance = 0
	source.Parent = nil

	queue := queues.NewQueue()
	queue.Enqueue(source)

	for !queue.Empty() {
		v := queue.Dequeue().(*graphs.Vertex)
		for _, adjVertex := range v.AdjList {
			if adjVertex.Color == colors.White {
				if adjVertex == end {
					return true
				}
				adjVertex.Color = colors.Gray
				adjVertex.Parent = v
				adjVertex.Distance = v.Distance + 1
				queue.Enqueue(adjVertex)
			}
		}
		v.Color = colors.Black
	}

	return false
}
