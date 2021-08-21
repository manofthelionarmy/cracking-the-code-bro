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
