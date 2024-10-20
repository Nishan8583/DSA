/*
4
WLLW
1 2 3 4
LWWL
6 5 3 2

1. A possible rival game is if one has L and another has W
2. The W team has more goal than L team
3. If rival game, get previous optimal game score and to current optimal
4. return max of either
*/
use std::{
    cmp::max,
    f32::consts::PI,
    io::{self, BufRead, StdinLock},
    process::exit,
    str::Lines,
};

fn main() {
    let mut i = 0;
    let mut n = 0;
    let mut result = 0;

    let ioreader = io::stdin();
    let mut reader = ioreader.lock().lines();

    let num_game = reader
        .next() // get the next line from stdin
        .unwrap_or(Ok(String::from("0"))) // handle the first Option type
        .unwrap_or(String::from("0")) // handle the Result type
        .trim() // trim spaces
        .parse::<usize>() // parse as usize
        .unwrap_or(0);

    let size = num_game + 1;
    let mut memo = vec![vec![-1; size]; size];

    let mut goose_outcome = vec!['R'; num_game + 1];
    let mut hawk_games = vec!['R'; num_game + 1];

    update_games(&mut reader, num_game, &mut goose_outcome);
    let goose_scores = update_score(&mut reader);

    update_games(&mut reader, num_game, &mut hawk_games);
    let hawk_score = update_score(&mut reader);

    //println!("Final goose_game {:?}", goose_outcome);
    //println!("Final socre {:?}", goose_scores);

    //println!("Final hawk_game {:?}", hawk_games);
    //println!("Final hawk_Score {:?}", hawk_score);

    let result = solve(
        &goose_outcome,
        &hawk_games,
        &goose_scores,
        &hawk_score,
        num_game,
        num_game,
        &mut memo,
    );

    println!("RESULT :{}", result);
}
// Lines<StdinLock<'static>>
// Any iterator that can yeild <Item = Result<String, io::Error>> is passed into this function
fn update_games<T>(reader: &mut T, num_game: usize, games: &mut Vec<char>)
where
    T: Iterator<Item = Result<String, io::Error>>,
{
    let goose_games = reader
        .next() // get the input
        .unwrap_or(Ok(String::from(""))) // unwrap outer Option
        .unwrap_or(String::from("")); // uinwrap inner result

    if goose_games.len() != num_game {
        /*println!(
            "Number of goose games are not enough. Expected {} got {}",
            num_game,
            goose_games.len()
        );*/
    }

    let mut char_iter = goose_games.chars();
    for i in 1..num_game + 1 {
        //println!("Inserting in {}", i);
        games[i] = char_iter.next().unwrap_or('R');
    }
}

fn update_score<T>(reader: &mut T) -> Vec<i32>
where
    T: Iterator<Item = Result<String, io::Error>>,
{
    let mut scores: Vec<i32> = reader
        .next()
        .unwrap_or(Ok(String::from("")))
        .unwrap_or(String::from(""))
        .split_whitespace()
        .filter_map(|x| x.parse::<i32>().ok())
        .collect();

    scores.insert(0, -1);
    return scores;
}
fn solve(
    outcome1: &Vec<char>,
    outcome2: &Vec<char>,
    goals1: &Vec<i32>,
    goals2: &Vec<i32>,
    i: usize,
    j: usize,
    memo: &mut Vec<Vec<i32>>,
) -> i32 {
    if memo[i][j] != -1 {
        return memo[i][j];
    }
    // if the goals we are considering in this subproblem is both 0, then return 0
    if i == 0 || j == 0 {
        memo[i][j] = 0;
        return 0;
    }

    let (mut first, mut second, mut third, mut fourth) = (0, 0, 0, 0);

    // to be a rival game, one of them muust have won and the other lost (both won or lost means different game)
    // and the winner must have had higher goal count, else does not make sense
    if ((outcome1[i] == 'W' && outcome2[j] == 'L') && (goals1[i] > goals2[j]))
        || ((outcome1[i] == 'L' && outcome2[j] == 'W') && (goals1[i] < goals2[j]))
    {
        // get the maximum goal in the rival game plus current goal
        /*println!(
            "Found rival game outcome1={} outcome2={} goal1={} goal2={} at i={} and j={}",
            outcome1[i], outcome2[j], goals1[i], goals2[j], i, j
        );*/
        first =
            solve(outcome1, outcome2, goals1, goals2, i - 1, j - 1, memo) + goals1[i] + goals2[j];
        //println!("Optimal {}", first);
    } else {
        first = 0;
    }
    second = solve(outcome1, outcome2, goals1, goals2, i - 1, j - 1, memo); // do not consider the current goal
    third = solve(outcome1, outcome2, goals1, goals2, i - 1, j, memo); // consider previous goal for outcome1
    fourth = solve(outcome1, outcome2, goals1, goals2, i, j - 1, memo); // consider previous goal for outcome 2

    /*println!(
        "First {} second {} third {} fourth {}",
        first, second, third, fourth
    );*/
    memo[i][j] = max(first, max(second, max(third, fourth)));
    //println!("Optimal for {} {} is {}", i, i, memo[i][j]);
    return memo[i][j];
}

#[cfg(test)]
mod test {
    use crate::solve;

    #[test]
    fn test_solve() {
        let mut memo = vec![vec![-1; 5]; 5];
        let max_goals = solve(
            &vec!['R', 'W', 'L', 'L', 'W'],
            &vec!['R', 'L', 'W', 'W', 'L'],
            &vec![-1, 1, 2, 3, 4],
            &vec![-2, 6, 5, 3, 2],
            4,
            4,
            &mut memo,
        );
        assert_eq!(14, max_goals);

        let mut memo = vec![vec![-1; 2]; 2];
        let max_goals = solve(
            &vec!['R', 'W'],
            &vec!['R', 'L'],
            &vec![-1, 1],
            &vec![-2, 6],
            1,
            1,
            &mut memo,
        );
        assert_eq!(0, max_goals);
    }
}
