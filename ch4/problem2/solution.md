# Solution

'To create a tree of minimal height, we need to match the number
of nodes in the left subtree to the number of nodes in the right
subtree as much as possible.'

- Cracking the Coding Interview, p. 242

# Initial Analysis (Not coding book's solution)

In this solution, they use an array to represent the binary search tree.


What's the difference between this array and creating a abstract data types?

Intuitively, (meaning what I think) is that the solution with the array takes up less memory. 

And the solution calls for creating a binary search tree with minimal height, meaning each right and left subtree
must have the same number of nodes. I'm pretty sure if you follow the algorithm, they will. I don't know in which case 
they will not have the same number of nodes... The input stated in the problem is sorted, increasing order; so each subtree
will have the same number of nodes. The binary search tree property will always ensure this. 

# New Analysis, (After coding book's solution)

Interestingly, we are doing a preorder traversal to create the tree. Makes sense, since we want to create the node
and fill in the left and right subtree. We are also doing divide-and-conquer approach by fragmenting the array in half,
and perpetually fragmenting each subarray in half until the fragments (subarrays) are not divisible. The mid-position is 
the right successor! Very interesting. And the root starts at the middle of the array. I can visually see how this ensures
that the tree will have minimal height (equal number of nodes in left and right subtree). Which leads to my comparison and
contrast between my solution and the book's solution.

# Compare and Contrast

## Compare:

Similarly, we are constructing the tree by filling in the left and right subtree. 

## Contrast:

My solution takes into consideration of filling the parent of each node. The book's solution doesn't.
However, in the book's defense, the question did ask for solely creating a binary search tree with minimal height
and didn't say anything about ensuring the parents are filled correctly.

The solution creates the root with value dead-center in the array. My solution creates the root at the first entry of 
the array! So, it makes me wonder if there's a difference between creating your roots with different values. 

I wonder if there's a way to prove that if you start your root with the value in the middle of your sorted, increasing
array, you will always have a binary tree with minimum height (equal number of nodes in right and left subtree).

The height of a tree is actually the longest path from the root to any leaf. Maximum height is the longest path. Minimum height is 
the shortest path from the root. Therefore, in my solution, because the input array is in sorted, increasing order and because I began
inserting nodes sequentially from the input array, the tree height will be of size len(input_array) 😬 (with no left subtree at all and 
everything being in the right subtree and will appear similar to a linked list)

Therefore, starting at the mid value in the input array will ensure we have about the equal number of nodes of left and right subtree
thanks to the binary search tree property.

# Things to work on 

Getting better at recursion. 

Maybe find a way to populate the parent field in the book's solution for creating a minimal-height binary search tree.

# Update: found a solution to update the parent field.

You have to use post order traversal. By time you have hit nil for both the left and right subtree, that means you just 
visited an external node (or edge). Recursively take steps back until root.left or root.right are not nil (meaning you 
are no longer at an edge, but at an internal node). That means we can set the parent of nodes root.right and root.left.

The runtime to do this is O(n). (right?)

Because a tree is like a graph, there are a total of n-1 'edges' (not an external node but lines that connect each node). And we visit
n nodes in the process. Therefore, O(n + n -1), thus O(n). We visit every node at least twice, but we drop that constant, O(2 * n) -> O(n).

