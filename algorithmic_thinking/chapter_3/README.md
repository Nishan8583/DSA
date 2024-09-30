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
