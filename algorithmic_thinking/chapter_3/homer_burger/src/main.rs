/*
We assume that the fulle time can be consumed by eating burger m, and so we keep on going consuming "m" if time "t" is enough, if not enought mark it as -1
Same for n
In final if both are -1, we could not solve them, so we keep on trying for t-1, t-2, until we find a time where homer can spend all the time eating burgers,
rest of time we eat beer.
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
            let v = solve(m, n, t);
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
