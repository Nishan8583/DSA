# Binary Search
- Divide search space in half
- Not just sorted array
- In the book example, feed ant array, Max liters we could use was 2B, we start with 2B, then half of that and so on, until we find something that does not feed the ant.
- When to use Binary Search
    1. **Hard Optimality and easy feasibility**: Like for ants, we did not know what the optimal solution was (did not know what minimum water in liter was), but we knew 20 liter would be enough and 10 would not be.
    2. **Infeasible-feasble split**: We can divide the border between feasible and infeasable space. Like in ant, we knew that anything above 20, could be feasble, and below 20 might not be (which we will keep dividing till we find).
- Use it if it is easiert to test feasibility than to find optimimal solution
- I see in books, places where any solution might work,there are multiple numbers that might work, so we try each, then divide space between feasable and non feasible, this division must be possible to use binary search.
- It does not have to be an array of increasing/decreasing values, just that after a certian operation half of it can be safely discarded, like in `cave dorrs` problem, we passed an array, with half switch turned on, if the door was open, it means the switch is within that range, if not, select another range and so on