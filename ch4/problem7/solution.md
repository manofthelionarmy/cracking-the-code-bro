# Build Order, book solution

Looks like I need to learn something new. We can use DFS. If we just went about this and traversed via DFS, we would always have a correct build order, but it won't be deterministic. 

In other words, there would be different build orders and they are still correct. 

The only time if it's impossible to get a generated build order if the graph is cyclical. We can tell if a graph is cyclical if we are trying to research a path that's being actively searched. 

so:

dfsHelper(g Graph, source *Vertex) bool {
    if source.Color == colors.Gray {
        return false
      }
  }

I beleive we can have a deterministic order in which we traverse via DFS and that order is determined via topological sort. 

TODO: study topological sort
