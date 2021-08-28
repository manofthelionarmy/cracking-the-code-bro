# My solution

I chose to create an undirected graph, instead of a directed graph. I don't have it implemented currently, so I wanted to save time.
BFS and DFS should work the same for both kinds of graphs.

I did some TDD and added some edges. I then did BFS so that a BFS tree can be produced. I then took a source and end vertex from the list
of vertices in the graph and back peddled through the parents of end vertex. If there is a path from the source to end vertex, then the parent
of endvertex's parent (an so forth) has to be the source vertex. If that's not the case, then the end vertex is not on the path.

# Book's solution

The solution was very interesting. It was to peform the BFS and traverse through the graph (using the search to traverse through the graph was interesting). You then check along the way if the end vertex is in the adjacency list of your source vertex.

The contrast from my solution to the book's solution is that I check the parent's of the endvertex. Looking at the book's solution, I could have gotten 
an equivalent result from checking the source's adjacency list and see if the end vertex is in there.
