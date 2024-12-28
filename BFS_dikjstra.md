## Breadth First search
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
| 1 | done  | 0 |
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
