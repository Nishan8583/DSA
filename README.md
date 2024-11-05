# DSA


# Memoization Dynamic programming
1. Problem used in book is UVa 10465. Read the problem pdf first.
2. Optimization problem, where we choose best solution out of all possible solution.
3. The suppostion we make to solve optimiztaion problem
  - **We suppose that homer can fill exactly t minutes with a burger**, if we are wrong, we assume that we can fill *t-1* minutes with burger, if that does not work then we check with *t-2* and so on.
  - The burger can be either **m** or **n**
  - If homer can spend entire *t* minutes eating burger, either m or n, then we can also assume that homer can also spend **t-m** or **t-n** minutes by exactly eating either **m** or **n** burgers.
  - But we are not sure, if homer can fill the entire time *t* or *t-1* and so on, with burgers, and we do not need to, we find out by solving each problem.
  - We need to check if homer can fill *t* by eating *m* or *n*, but before that we need to find if he can fill *t-m* or *t-n* with either *m* or *n*, so we need to solve that first
  - But before we solve that we need to check if *(t-m)-m* or *(t-n)-n* can be filled with burger or no, so we solve that subprobme, and thus subproblem.
  - If it turns out we can not then we search if it can be for *t-1* and for *t-2*, and remaining time we dring bear.
4. Ok so how do we do that in code?
5. We solve recursively,
  - for both t-m and t-n, and keep on going if in any branch we reach 0 it means we could spend entire time eating bruger,
  - if not we solve it again for t-1, and t-2 and so on,
  - until we find any time that was able to spend eating burger fully, and rest time we eat beer.


![flow](logic)

## Memoization
- In figure we can see that a lot of problems have been solved more than once, memoize it. Keep the record in array

## Dynamic programming
- Solve from t=0, and for each value get t-n which is recorded an an array.
- In DP, we reorganize the solution in such a way, the subproblem is solved by the time we need it.
- In memoziation we store the solution.
## Steps on
1. Check if problem can be divided in subproblems for the optimal solution.
2. Create a recursive solution.
3. Check if there is overlapping subproblems.
4. For memoization, record the solution. Top down, because we see from top bigger solution and go to the buttom.
5. For DP, solve the smallest subproblem first, record its solution, and then grudually go to bigger problem and then get recorded solution as required. Bottom up, because we solve the bottom most solution first.


## When to use them?
- When question is asking for minimum or maximum of something, we try to find minimum or maximum of smaller problem, like maximum burgers u can eat in time `t`, we try to find maximum we can eat in `t-1` time with both `m` and `n` burgers. Or for house robber problem, we find max we can steal from smaller subarray, and then we include bigger subarray in the context.
- To solve smaller subarray we use recursion.
- If we keep on sovling same problem again and again, memoize.
- We could intead start by solving smaller subproblems first.



# Breadth First search
- Search full range of current avaiable, and then move to next.
- Use it whenever u want to find minimum distance from one location to another location.
- Use it also to find distance from one position to all position
- Visualize, create graph, each nodes like (chess box, index in a rope), connection draw as edge.
- How to distinguish between DP? Presence of cycle is definetly a graph. Like (in apple problem, no way to get to higher number of apple so no BFS)
- For optimization, decreace the number of BFS calls, and decrease the number of edges (maybe use 2 arrays to save state rather than the number of connection)
- To achieve BFS. use function call to go to all available current position, or use queue?

# Dijkstra's Algorithm
- Usefule to find shotest path from point `a` to point `b` in a weighted graph.
- Consider the problem [Mice and Maze](https://vjudge.net/problem/UVA-1112).
- The question asks us to find the total number of mice can reach the exit cell within given time limit.
- So for each cell, we find the shortest path to that exit cell.
- The edges will be something like this in go, with indices representing start position, end position and the time length (weight of the edge).
- The code we refeerence to is in [chapter_6](https://github.com/Nishan8583/DSA/blob/master/algorithmic_thinking/chapter_6/mice_maze/main.go).
```go
	  {1, 2, 12},
		{1, 3, 6},
		{2, 1, 26},
		{1, 4, 45},
		{1, 5, 7},
		{3, 2, 2},
		{2, 4, 9},
		{4, 3, 8},
		{5, 2, 21},
```
- Visualized looks like this
![image](https://github.com/user-attachments/assets/f793bf16-6971-47e0-88de-1c59c5f77223)

- For each starting position we find the shotest path to all the nodes.
```go
total := 0
		for i := 0; i <= num_cells; i++ {
			min_time := find_time(adj_list, num_cells, i, exit_cell)
			if min_time >= 0 && min_time <= time_limit {
				total++
			}
		}
```
- Example for node 1:
- We keep track of `done` (the shortest path to this cell has been found), and `min_time` (minumum time to get to the cell).
- From 1, the visualization table will be something like this

| Node       | Done       | min_index       |
| -------------- | -------------- | -------------- |
| 1 | false | 0 |
| 2 | false | -1 |
| 3 | false | -1 |
| 4 | false | -1 |
| 5 | false | -1 |
```go
	for i := 1; i <= num_cells; i++ {
		done[i] = false
		min_times[i] = -1
	}
	min_times[from_cell] = 0
```
- 0 time to get to itself and so its done
- Now update the index with direct connection values.
| Node       | Done       | min_index       |
| -------------- | -------------- | -------------- |
| 1 | done | 0 |
| 2 | false | 12 |
| 3 | false | 6 |
| 4 | false | 45 |
| 5 | false | 7 |

```go
	for i := 0; i < num_cells; i++ {
		min_time = -1
		found = false

		for j := 1; j <= num_cells; j++ {
			// The cell has to be not done and its min_times has to be atleast greater than or equals to 0
			if !done[j] && min_times[j] >= 0 {
				// if min_time has not be set yet or we found some new time which is less than  previous min time
				if min_time == -1 || min_times[j] < min_time {
					min_time = min_times[j]
					min_time_index = j
					found = true
				}
			}
		}
  ---code---snippet--
	}
```
- THe most minimum value here is `6` to get to node 3.
- Since it is already the most min value, even if there is another path to it, it would be automatically be at least time `6`, so mark it as done.
| Node       | Done       | min_index       |
| -------------- | -------------- | -------------- |
| 1 | done | 0 |
| 2 | false | 12 |
| 3 | done | 6 |
| 4 | false | 45 |
| 5 | false | 7 |

- Now we keep on repeating, check the most min time besides the `done` nodes.
- Its node 5
- Can we see if we can reach anything else from node 5, which might produce shorter value, yes, we can reach node 2, with total final value of 8
| Node       | Done       | min_index       |
| -------------- | -------------- | -------------- |
| 1 | done | 0 |
| 2 | false | 12 |
| 3 | done | 6 |
| 4 | false | 45 |
| 5 | false | 7 |


```go
	for i := 0; i < num_cells; i++ {
		min_time = -1
		found = false
----code -----

		e = adj_list[min_time_index]
		for e != nil {
			old_time = min_times[e.to_cell]
			if old_time == -1 || old_time > min_time+e.length {
				min_times[e.to_cell] = min_time + e.length
			}
			e = e.next
		}
	}
```
- We keep on repeating until we have completed all done node.

# Binary Search
- Divide search space in half
- Not just sorted array
- In the book example, feed ant array, Max liters we could use was 2B, we start with 2B, then half of that and so on, until we find something that does not feed the ant.
- When to use Binary Search
    1. **Hard Optimality and easy feasibility**: Like for ants, we did not know what the optimal solution was (did not know what minimum water in liter was), but we knew 20 liter would be enough and 10 would not be.
    2. **Infeasible-feasble split**: We can divide the border between feasible and infeasable space. Like in ant, we knew that anything above 20, could be feasble, and below 20 might not be (which we will keep dividing till we find).
