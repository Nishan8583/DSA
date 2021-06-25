fn binary_search(A: &[i32], x: i32) -> Option<usize> {
    let mut start = 0;
    let mut end = A.len()-1;

    while start <= end {
        let mid = (start+end)/2;
        println!("{} {} {}",start,end,mid);
        if A[mid] == x {
            return Some(mid);
        } else if x < A[mid] {
            end = mid - 1;
        } else {
            start = mid + 1;
        }
    }

    None
}

fn binary_search_iter(a: &[i32], mut start: usize, mut end: usize, x: i32) -> Option<usize> {
    // out of bound ?
    if start >= end {
        return None;
    }

    let mid = (start+end)/2;

    if a[mid] == x {
        return Some(mid);
    } else if x > a[mid] {
        start = mid + 1;
    } else {
        end = mid - 1;
    }

    return binary_search_iter(a,start,end,x);
}
fn main() {
    let array = [1,2,3,4,5,6,7,8,9];
    let mut b = array.len();
    let i = match binary_search_iter(&array,0,b,9) {
        Some(v) => v,
        None => panic!("value not found"),
    };
    println!("Position is at {}",i);
}