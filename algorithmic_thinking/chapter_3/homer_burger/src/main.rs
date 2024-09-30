/*
We assume that the fulle time can be consumed by eating burger m, and so we keep on going consuming "m" if time "t" is enough, if not enought mark it as -1
Same for n
In final if both are -1, we could not solve them, so we keep on trying for t-1, t-2, until we find a time where homer can spend all the time eating burgers,
rest of time we eat beer.

OR I can think of it as a decision tree
We canculate for a time when we eat burger "m" and when eat "n", and check substract the time on each call with t-m and t-n respectively.
And then we decide which condition gives us non negative value, Memoize ti

For DP solution, we start with at each time, from 0, how many burgers can it eat, either m or n, whichever is non zero
*/

use std::{
    cmp::max,
    io::{self, BufRead},
};

fn main() {
    let stdin = io::stdin();

    let mut lines = stdin.lock().lines(); // read input from stdin line by line

    while let Some(Ok(line)) = lines.next() {
        // while reading line gives Some(Ok()) value
        let mut values = line
            .split_whitespace()
            .filter_map(|s| s.parse::<i32>().ok());

        if let (Some(m), Some(n), Some(t)) = (values.next(), values.next(), values.next()) {
            //let v = solve(m, n, t);
            solve_dp(m, n, t);
        } else {
            break;
        }
    }
}

fn solve_t(m: i32, n: i32, t: i32, calls: &mut i32, set: &mut [i32]) -> i32 {
    *calls += 1;
    // base case
    if t == 0 {
        return 0;
    }

    let mut first = 0;
    let mut second = 0;

    if set[t as usize] != -2 {
        return set[t as usize];
    }

    // first case, assumes that we can fill exactly t minutes by eating m burgers
    if t >= m {
        first = solve_t(m, n, t - m, calls, set);
        set[(t - m) as usize] = first;
    } else {
        first = -1;
    }

    // second case, assumes that we can fill exactly t minutes by eating n burgers
    if t >= n {
        second = solve_t(m, n, t - n, calls, set);
        set[(t - n) as usize] = second;
    } else {
        second = -1;
    }

    if first == -1 && second == -1 {
        return -1;
    }

    let ret = max(first, second) + 1;
    return ret;
}

fn solve(m: i32, n: i32, t: i32) -> i32 {
    let mut i = 0;
    let mut calls = 0;

    let mut memo = [-2; 10000];
    let mut result = solve_t(m, n, t, &mut calls, &mut memo);

    if result >= 0 {
        println!("{}", result);
        return result;
    } else {
        i = t - 1;
        result = solve_t(m, n, i, &mut calls, &mut memo);
        while result == -1 {
            i -= 1;
            result = solve_t(m, n, i, &mut calls, &mut memo);
        }
        println!("{} {}", result, t - i);
    }

    return -1;
}
fn solve_dp(m: i32, n: i32, t: i32) {
    // for explaination lets assume, m = 4, n =9, and t = 15
    let mut dp = [-1; 10000];

    dp[0] = 0; // in time t = 0, there is exactly 0 burgers left because we can not eat anything
    for i in 1..t + 1 {
        let mut first = -1;
        if i >= m {
            // till 1-3, the values will be -1
            // |    0   |   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |
            // for i = 4, i-m=4-4=0,dp[0],0
            // |    0   |   -1  |   -1  |   -1  |
            // for m = 8
            // |    0   |   -1  |   -1  |   -1  |
            //println!("Calculate for i={} m={}", i, m);
            first = dp[(i - m) as usize]; // when i = 4, its going to be 4-4, dp[0] is 0
        }

        let mut second = -1;
        if i >= n {
            //println!("Calculate for i={} n={} difference {}", i, n, i - n);
            second = dp[(i - n) as usize];
        }

        if first == -1 && second == -1 {
            dp[i as usize] = -1;
            println!("Setting index {} as {}", i, -1);
        } else {
            dp[i as usize] = max(first, second) + 1; // here is where we add this
            println!("Non -1 for {} {}", i as usize, max(first, second) + 1);
        }
    }

    if dp[t as usize] >= 0 {
        println!("{}", dp[t as usize]);
    }

    let mut result = dp[t as usize];
    if result >= 0 {
        println!("{}", result);
    } else {
        let mut i = t - 1;
        result = dp[i as usize];
        while (result == -1) {
            i -= 1;
            println!("Checking for {}", i);
            if i < 0 {
                println!("No solution found, index {}", i);
                break;
            }
            result = dp[i as usize];
            println!("value {}", result);
        }
        println!("{} {}", result, t - i);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve_t() {
        let mut calls = 0;

        let mut memo = [-2; 10000];
        assert_eq!(3, solve_t(4, 9, 22, &mut calls, &mut memo));
        println!("Number of calls made {}", calls);
        calls = 0;

        let mut memo = [-2; 10000];
        assert_eq!(5, solve_t(4, 5, 20, &mut calls, &mut memo));

        println!("Number of calls made {}", calls);
        calls = 0;

        let mut memo = [-2; 10000];
        assert_eq!(-1, solve_t(4, 9, 15, &mut calls, &mut memo));

        println!("Number of calls made {}", calls);

        calls = 0;

        let mut memo = [-2; 10000];
        assert_eq!(44, solve_t(4, 2, 88, &mut calls, &mut memo));

        println!("Number of calls made {}", calls);

        calls = 0;

        let mut memo = [-2; 10000];
        assert_eq!(45, solve_t(4, 2, 90, &mut calls, &mut memo));

        println!("Number of calls made {}", calls);
    }

    #[test]
    fn test_solve() {
        assert_eq!(3, solve(4, 9, 22));

        assert_eq!(44, solve(4, 2, 88));
    }
}
