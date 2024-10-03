/*

To get cost of n apples We get the minumum price of n-1 apples
TO get price of apple, we check both using unit price, or by applying scheme,
whicever is lower, use that
Also buying more apple might be cheaper, at max can buy 200 apples, so loop till 200

    We get the price for all cases
    1. Get minimum price of n-1 apples, add unit price, treat it as best
    2. Loop through all the schemes, check if the scheme actually applied to the number of apples
    3. If it does, deecrease the number of apples with the one in the scheme, and calcualte the beest price for the rest of the apples
    4. check if the new price is better than the previous one, if so update it.


    Example:
    1.75 2
    4 3.00
    2 2.00
    3
    Here we, unti price for apples is 1.75, we have 2 schemes
    4 apples for 3 dollars
    2 apples for 2 dollars
    We need to buy 3 apples

    num_apples_per_scheme=[4,2]
    price_per_scheme=[3,2]
    num_schemes=2
    unit_price=1.75
    num_apples=3

    So we want to buy 3 apples
    1. we reach the bottom most first, 0 apples, returns 0
    Then we go up one above, 1 apples cost + 0 apple cost //   let mut best = result + unit_price;
    2. Now check if any scheme available for 1 apple, no so go up, return 1.75
    3. Now for 2 apple, cost of current unit apple + cost of 2-1 apple = 1.75+1.75=3.5
    4. Check if scheme present, yes, because 2, substract number of apples from scheme 2-2. 0, find min price of 0 apple which is 0,
    5. Add result (cost of 0 apple) + priceSCheme for 2(2.00), compare with previous best, if minimum return
    6. One level higher for 3, cost of current price n + n-3, cost of 3rd apple + cost of two apples = 2+1.75=3.75
    7. Chck if scheme, yes, scheme 2 matches, deduct 2, i.e. 3-2. get minimum cost for 1 apple,
    8. add result (cost of 1 apple) with price schemee for 2 apple. check if lower, if so return
    9. Now for 4, cost of 4th apple + minimum cost of (4-1) 3 apple. =5.5
    10. Check if scheme, yes scheme for 4, 4-4 = 0, cost of 0 apple + scheme price = 3. yes lower so return

*/

use std::io::{stdin, BufRead};

fn main() {
    let mut reader = stdin().lock().lines();

    println!("Enter apple booiz");
    while let Some(Ok(line)) = reader.next() {
        // after done with the filter_map, it will be a iterator as well
        let mut values = line
            .split_whitespace() // split by whitespaces, returns iterator
            .filter_map(|x| x.parse::<f32>().ok()); // on each element of white spaces, run the following function i.e. parse into i32, and return as tuples

        let (unit_price, test_cases) = (values.next().unwrap(), values.next().unwrap());
        let mut schemes_apple_count = vec![];
        let mut schemes_apple_cost = vec![];
        let mut count = test_cases as usize;

        while count > 0 {
            let test_line = reader.next().unwrap().ok().unwrap();
            let mut test_values = test_line
                .split_whitespace()
                .filter_map(|x| x.parse::<f32>().ok());

            let (apple_count, scheme_price) =
                (test_values.next().unwrap(), test_values.next().unwrap());

            schemes_apple_count.push(apple_count);
            schemes_apple_cost.push(scheme_price);
            count -= 1;
        }
        println!(
            "Unit Price={} Test cases={} Schemes_apples={:?} Schemes_price={:?}",
            unit_price, test_cases, schemes_apple_count, schemes_apple_cost
        );

        println!("Enter apple booiz");
    }

    println!("Hello, world!");
}

fn solve_k(
    num_apples_per_scheme: Vec<i32>, // array number of apples per scheme, if there are scheme for 3 and 2 apples, it will be [3,2]
    price_per_scheme: Vec<f32>, // price for each scheme, example, if first scheme, i.e. for 3 apples is 4 and for 2 apples is 2.50. array is [4.00,2.50]
    num_schemes: i32, // total number of schemes, I guess we could have gotten that from first arguement
    unit_price: f32,  // price of one apple
    num_apples: i32,  // the number of apples we want to buy
    memo: &mut [f32; 200],
) -> f32 {
    //println!("Searching for {}", num_apples);
    // base case, it takes 0 money to buy 0 apples
    if num_apples == 0 {
        return 0.0;
    }

    let temp_index = num_apples as usize;
    if memo[temp_index] != -1.0 {
        //println!("Found in cache for {}", memo[temp_index]);
        return memo[temp_index];
    }
    let mut result = solve_k(
        num_apples_per_scheme.clone(),
        price_per_scheme.clone(),
        num_schemes,
        unit_price,
        num_apples - 1,
        memo,
    );
    let mut best = result + unit_price;
    //println!("**************initially for {} is {}", num_apples, best);
    for i in 0..num_schemes {
        let ic = i as usize;
        //println!("Checking for scheme {}", num_apples_per_scheme[ic]);
        if (num_apples - num_apples_per_scheme[ic]) >= 0 {
            //println!("Solving for {} {}", i, num_apples_per_scheme[ic]);

            result = solve_k(
                num_apples_per_scheme.clone(),
                price_per_scheme.clone(),
                num_schemes,
                unit_price,
                num_apples - num_apples_per_scheme[ic], // so i am decreasing the number of apples to buy minus the number of apples in schemes
                //so lets say we have scheme of 2 for 2.00 dollars ,then we decrease number of apples to buy by 2
                // which returns 0
                // and then in the min() step we add result wiht price of this schemee
                memo,
            );

            best = min(result + price_per_scheme[ic], best);
        } else {
            println!("Notsuitable for scheme {}", num_apples_per_scheme[ic]);
        }
    }

    println!(
        "Appending in memo {} for {} appending in memo",
        best, temp_index
    );
    memo[temp_index] = best;
    return best;
}

fn min(a: f32, b: f32) -> f32 {
    if a <= b {
        return a;
    }
    return b;
}
const size: i32 = 200;

fn solve(
    num_apples_per_scheme: Vec<i32>, // array number of apples per scheme, if there are scheme for 3 and 2 apples, it will be [3,2]
    price_per_scheme: Vec<f32>, // price for each scheme, example, if first scheme, i.e. for 3 apples is 4 and for 2 apples is 2.50. array is [4.00,2.50]
    num_schemes: i32, // total number of schemes, I guess we could have gotten that from first arguement
    unit_price: f32,  // price of one apple
    num_apples: i32,  // the number of apples we want to buy)
) -> f32 {
    let mut memo: [f32; 200] = [-1.0; 200];

    let mut best = solve_k(
        num_apples_per_scheme.clone(),
        price_per_scheme.clone(),
        num_schemes,
        unit_price,
        num_apples,
        &mut memo,
    );
    println!("Initial best is {}", best);
    for i in num_apples..size {
        best = min(
            best,
            solve_k(
                num_apples_per_scheme.clone(),
                price_per_scheme.clone(),
                num_schemes,
                unit_price,
                i,
                &mut memo,
            ),
        );
        println!("Loop best for {} is {}", i, best);
    }
    return best;
}
#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_solve_k() {
        let mut s = [-1.0; 200];
        let best = solve_k(vec![4, 2], vec![3.00, 2.00], 2, 1.75, 4, &mut s);
        assert_eq!(3.00, best);

        let best = solve(vec![4, 2], vec![3.00, 2.00], 2, 1.75, 4);
        assert_eq!(3.00, best);
    }
}
