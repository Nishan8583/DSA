# DSA

## DP
#### When to use ?
- When question is asking for minimum or maximum of something, we try to find minimum or maximum of smaller problem, like maximum burgers u can eat in time `t`, we try to find maximum we can eat in `t-1` time with both `m` and `n` burgers. Or for house robber problem, we find max we can steal from smaller subarray, and then we include bigger subarray in the context.
- To solve smaller subarray we use recursion.
- If we keep on sovling same problem again and again, memoize.
- We could intead start by solving smaller subproblems first.
[DP](./DP.md)
## BFS and Dikstra
#### When to use BFS?
- Use it whenever u want to find minimum distance from one location to another location.
- Use it also to find distance from one position to all position
- How to distinguish between DP? Presence of cycle is definetly a graph. Like (in apple problem, no way to get to higher number of apple so no BFS
#### When to use Dikstra?
- Usefule to find shotest path from point `a` to point `b` in a weighted graph.
[BFS_dikstra](./BFS_dikjstra.md)

## Binary Search
#### When to use Binary Search ?
    1. **Hard Optimality and easy feasibility**: Like for ants, we did not know what the optimal solution was (did not know what minimum water in liter was), but we knew 20 liter would be enough and 10 would not be.
    2. **Infeasible-feasble split**: We can divide the border between feasible and infeasable space. Like in ant, we knew that anything above 20, could be feasble, and below 20 might not be (which we will keep dividing till we find).
- Use it if it is easiert to test feasibility than to find optimimal solution
- I see in books, places where any solution might work,there are multiple numbers that might work, so we try each, then divide space between feasable and non feasible, this division must be possible to use binary search.
[binary_serach](./binary_search.md)

# Heap
#### When to use heaps?
- Use it when the question requires to get maximum and minimum value of some sort
[Heap](./Heaps.md)

# Segment Trees
####  When to use Segment Trees ?
- When we want some max/min value within certain range.
[SegmentTree](./segment_trees.md)