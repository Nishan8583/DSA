use std::cmp::min;

fn main() {
    println!("Hello, world!");
}

fn solve(
    costs: &Vec<i32>,
    num_squares: usize,
    current_square: i32,
    jump_distance_to_current_square: i32,
    memo: &mut Vec<Vec<i32>>,
) -> i32 {
    // if we found something in the memo
    let cur_index = current_square as usize;
    let cur_num_jumps = jump_distance_to_current_square as usize;
    //println!(" current cost index {}", costs[cur_index]);
    if let Some(square_memo) = memo.get(cur_index) {
        if let Some(cost) = square_memo.get(cur_num_jumps) {
            if *cost != -2 {
                println!("Found cahce for {} {}", cur_index, cur_num_jumps);
                return *cost;
            }
        }
    }

    // if we are in the final square, the cost to move to this square is always gonna be 0, since we are already here
    if cur_index == num_squares {
        println!(
            "---------------------final square, cahce it {} {}",
            cur_index, num_squares
        );
        println!("ACCESSING");
        memo[num_squares][cur_num_jumps] = 0;
        println!("ACcess done");
        return 0;
    }

    let mut first = -1;
    let mut second = -1;

    if (current_square + jump_distance_to_current_square + 1) <= num_squares as i32 {
        println!(
            "Checkin found jump={} num_sqare={}",
            (current_square + jump_distance_to_current_square + 1),
            num_squares as usize
        );
        println!(
            "Jumping right from index={} value={} with jump distance {}>>>>>",
            cur_index,
            costs[cur_index],
            jump_distance_to_current_square + 1,
        );
        first = solve(
            costs,
            num_squares,
            current_square + jump_distance_to_current_square + 1, // jumping to the right square
            jump_distance_to_current_square + 1,                  // increase the number of jumps
            memo,
        );
    }

    // if it is possible to jump left
    if current_square - jump_distance_to_current_square >= 1 {
        println!(
            "Jumping left from {} value={} <<<<<<<<<<<<<",
            cur_index, costs[cur_index]
        );
        second = solve(
            costs,
            num_squares,
            current_square - jump_distance_to_current_square, // jump left
            jump_distance_to_current_square,
            memo,
        );
    }

    if (first == -1) && (second == -1) {
        memo[cur_index][cur_num_jumps] = -1;
        println!("Min cost cound not be found  {} return -1", cur_index);
        return memo[cur_index][cur_num_jumps];
    } else if second == -1 {
        memo[cur_index][cur_num_jumps] = first + costs[cur_index + cur_num_jumps + 1];
        println!(
            "Min cost found when jumping right for {} cost= {} cause left gives -1",
            cur_index, memo[cur_index][cur_num_jumps]
        );
        return memo[cur_index][cur_num_jumps];
    } else if first == -1 {
        memo[cur_index][cur_num_jumps] = costs[cur_index - cur_num_jumps] + second;
        println!(
            "Min cost found when jumping left for {} cost= {} cause right gives -1",
            cur_index, memo[cur_index][cur_num_jumps]
        );
        return memo[cur_index][cur_num_jumps];
    }

    println!("Calculating min");
    memo[cur_index][cur_num_jumps] = min(
        costs[cur_index + cur_num_jumps + 1] + first,
        costs[cur_index - cur_num_jumps] + second,
    );
    println!(
        "Min cost found when jumping right for {} cost= {}",
        cur_index, memo[cur_index][cur_num_jumps]
    );
    return memo[cur_index][cur_num_jumps];
}

#[cfg(test)]
mod tests {
    use crate::solve;

    #[test]
    fn test_solve() {
        let size = 1000;
        let mut costs = vec![0; size + 1];
        costs[1] = 3;
        costs[2] = 5;
        costs[3] = 1;
        costs[4] = 9;
        costs[5] = 7;
        costs[6] = 2;
        costs[7] = 3;
        let mut memo: Vec<Vec<i32>> = vec![vec![-2; costs.len() + 1]; costs.len() + 1];
        println!("Memo {:?}, {}", memo, memo.len());
        let result = costs[2] + solve(&costs, 7, 2, 1, &mut memo);
        assert_eq!(15, result);
    }
}
