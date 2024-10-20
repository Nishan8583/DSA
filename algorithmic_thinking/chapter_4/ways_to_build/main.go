package main

import "fmt"

type pair struct {
	end_at_i int
	total    int
}

const (
	MAX_A = 1000
	MAX_B = 200
	MAX_K = 200
	MOD   = 1000000007
)

// i index for a
// j index for b
// k is number of substring we can use
func solve(a string, b string, i, j, k int, memo [MAX_A][MAX_B][MAX_K + 1]pair) pair {

	if memo[i][j][k].total != -1 {
		return memo[i][j][k]
	}

	var total, end_at_i = 0, 0
	// if at final index for b, and final substring we can use is 1
	if j == 0 && k == 1 {
		if a[i] != b[j] {
			if i == 0 {

				total = 0
			} else {
				total = solve(a, b, i-1, j, k, memo).total
			}
			memo[i][k][k] = pair{i, total}
		} else {
			if i == 0 {
				total = 1
			} else {
				total = 1 + solve(a, b, i-1, j, k, memo).total
			}
			memo[i][j][k] = pair{1, total}
		}
		return memo[i][j][k]
	}

	if (i == 0) || j == 0 || k == 0 {
		memo[i][j][k] = pair{0, 0}
		return memo[i][j][k]
	}

	if a[i] != b[j] {
		end_at_i = 0
	} else {
		end_at_i = solve(a, b, i-1, j-1, k-1, memo).total + solve(a, b, i-1, j-1, k, memo).end_at_i
		end_at_i = end_at_i % MOD
	}

	total = (end_at_i + solve(a, b, i-1, j, k, memo).total) % MOD
	memo[i][j][k] = pair{end_at_i, total}
	return memo[i][j][k]
}
func main() {
	a_length := 0
	b_length := 0
	num_substrings := 0

	a := ""
	b := ""

	memo := [MAX_A][MAX_B][MAX_K + 1]pair{}

	fmt.Scanf("%d%d%d", &a_length, &b_length, &num_substrings)
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	for i := 0; i < a_length; i++ {
		for j := 0; j < b_length; j++ {
			for k := 0; k <= num_substrings; k++ {
				memo[i][j][k] = pair{-1, -1}
			}
		}
	}
	result := solve(a, b, a_length-1, b_length-1, num_substrings, memo).total
	fmt.Printf("%d", result)

}
