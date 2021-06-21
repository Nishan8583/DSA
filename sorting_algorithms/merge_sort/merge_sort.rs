// mergint the left and right array into a
fn merge(l: Vec<i32>, r: Vec<i32>, mut a: Vec<i32>) -> Vec<i32> {
    
    // li is position for l array, ri array, and ai for a array
    let mut li = 0;
    let mut ri = 0;
    let mut ai = 0;

    while li < l.len() && ri < r.len() {
        if l[li] <= r[ri] {
            a[ai] = l[li];
            li = li + 1;
        }else{
            a[ai] = r[ri];
            ri = ri + 1;
        }
        ai = ai + 1;
    }
    a
}

// implements the logic of breaking down bigger vector into smaller ones
fn merge_sort(mut a:  Vec<i32>) -> Vec<i32>{
    
    // it is just a single element
    if a.len() < 2{
        return a
    }

    // breaking down into two parts
    let mid = a.len()/2;
    let mut left = vec![0;mid];
    let mut right = vec![0;a.len()-mid];

    // filling first part
    for i in 0..mid{
        left[i] = a[i];
    }

    // filling second part
    for i in mid..a.len(){
        right[i-mid] =  a[i];
    }

    // mergint the subparts first
    let mut left = merge_sort(left);
    let mut right = merge_sort(right);

    let merged_array = merge(left,right,a);
    return merged_array;
}

fn main() {
    let mut a = vec![2,4,1,6,8,5,3,7];
    let a = merge_sort(a);
    println!("{:?}",a);
}