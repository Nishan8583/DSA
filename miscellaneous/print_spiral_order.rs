fn spiral_order_clock_wise(a: &Vec<Vec<i32>>, m: usize, n: usize) {
    let mut t = 0;  // will point to the top most row of untraversed array
    let mut b = m;  // will point to the bottom most row of untraversed array
    let mut l = 0;  //will point to the leftmost column of untraversed array
    let mut r = n; // will point to the right most column of untraversed array

    // the direction to go to
    // 0 -> left to right
    // 1 -> top to bottom
    // 2 -> right to left
    // 3 -> bottom to up
    let mut dir = 0;

    while (t <= b && l <= r) {
        match dir {
            // left to right
            0 => {
                for i in l..r {
                    println!("left to right {}",a[t][i]);
                }
                t = t + 1;
                dir = 1;
            },
            // top to bottom
            1 => {
                for i in t..b {
                    println!("top to bottom {}",a[i][r-1]);
                }
                r = r - 1;
                dir = 2;
            },
            // right to left
            2 => {
                println!("right to left {} {} {}",dir,r,l);
                for i in (l..r).rev() {
                    println!("right to left {}",a[b-1][i]);
                }
                b = b - 1;
                dir = 3;
            },
            // bottom to top
            3 => {
                for i in (t..b).rev() {
                    println!("bottom to up {}",a[i][l]);
                }
                l = l + 1;
                dir = 0;
            },
            _ => {},
        };
    }
    
}

fn main() {
    let v = vec![
        vec![1,2,3,4],
        vec![5,6,7,8],
        vec![9,10,11,12],
        vec![13,14,15,16],
        ];
    spiral_order_clock_wise(&v,4,4);
}