fn fib_naive_recur(i: i32) -> i32 {

    if i == 0 {
        return 0;
    }

    if i == 1 || i == 2 {
        return 1;
    }

    return fib_naive_recur(i-1) + fib_naive_recur(i-2);
}

fn fib_recur_memo(n: usize, memo: &mut Vec<i32>) -> i32 {
    
    // if there is actual value in that index which is not -1,
    // then the value has been calculated
    match memo.get(n) {
        Some(v) => {
            if *v !=  -1{
                return *v;
            }
        },
        None => {}
    };

    let mut result = 0;

    if n == 1 || n == 2 {
        result = 1;
    } else {
        result = fib_recur_memo(n-1,memo) + fib_recur_memo(n - 2,memo);
    }
    // putting the value in the vector so that we do not have to calculate it again
    memo.insert(n,result);
    return memo[n];
}

fn fib_bottom_up(n: i64) -> i64 {
    if n == 0 {
        return 0;
    }

    if n == 1 || n == 2 {
        return 1;
    }

    let vecSize = (n+1) as usize;
    let mut bottom = vec![0;vecSize];
    bottom.insert(1,1);
    bottom.insert(2,1);

    for i in 2..n+1 {
        let index = i as usize;
        //println!("index {}",index);
        let value = bottom[index-1]+bottom[index-2];
        //println!("value to insert {}",value);
        bottom.insert(index,value);
    }

    return bottom[n as usize];
}

// floud warshall
// consider a graph of following type 
/*
         A   B   C   D
A        0   1   2   3
    
B    0   0   3   -1  7

C    2  8 ...

D    3
*/
fn all_pairs_shortest_path() {
    let mut arr = [
        [0,3,300,7],
        [8,0,2,300],
        [5,300,0,1],
        [4,300,300,0],
    ];
    println!("{:?}",arr);

    // k is the intermediate point
    for k in 0..arr.len() {
        for i in 0..arr.len() {
            for j in 0..arr.len(){
                let v = arr[i][k] + arr[k][j];
                if arr[i][j] > v {
                    arr[i][j] = v;
                }
            }
        }
    }

    println!("{:?} shortest path becomes",arr);
}
fn main() {
    //println!("{}",fib_naive_recur(40));
    //let mut memo = vec![-1;50+1];
    //println!("{}",fib_recur_memo(50,&mut memo));
    println!("for 0 {}",fib_bottom_up(0));
    println!("for 1 {}",fib_bottom_up(1));
    println!("for 2 {}",fib_bottom_up(2));
    println!("for 3 {}",fib_bottom_up(3));
    println!("for 5 {}",fib_bottom_up(5));
    println!("for 50 {}",fib_bottom_up(50));
    all_pairs_shortest_path();
}