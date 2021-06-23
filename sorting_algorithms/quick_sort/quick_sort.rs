fn partition(mut a: &mut [i32], start: usize, end:usize) -> usize{
    println!("start is {} and end is {}",start,end);
    let pivot = a[end];
    let mut p_index = start;

    for i in start..end {
        if a[i] <= pivot {
            let temp = a[i];
            a[i] = a[p_index];
            a[p_index] = temp;
            p_index = p_index + 1;
        }
    }
    
    
    let temp = a[p_index];
    a[p_index] = a[end];
    a[end] = temp;
    println!("SWapped array {:?}",a);
    return p_index;

}

/*
[7,2,1,6,5,3,4]
p_index = 0 = 7
i = 0,
7 is not less than 4 so no swapping

i = 1, p_index still same
2 is less than 2 so swapping 7 and 2, 
[2,7,1,6,5,3,4]

i = 2, p_index = 1
1 is less than 4 so again
[2,1,7,6,5,3,4], p_index = 2 now

i = 3, 6 is not less than 4 so no swapping,
i = 4, 5 is not less
i = 5, 3 is less so swap with p_index
[2,1,3,6,5,7,4] p_idnex = 3 now

the below step will not be reached, cause the loop ends just before it
i = 6, 4 is less than and equals so swapping
[2,1,3,4,5,7,6] p_index = 4

so, the final array in loop becomes
[2,1,3,6,5,7,4]

outsied loop will make this
[2,1,3,4,5,7,6]
3 is returned
*/

fn quick_sort(mut a: &mut [i32], start: usize, end: usize) {
    if start >= end {
        return
    }
    let p_index = partition(&mut a, start, end);
    println!("returned value {}",p_index);
    if p_index != 0  {
        quick_sort(&mut a,start, p_index-1);

    }
    quick_sort(&mut a, p_index+1,end);
}

fn main() {
    let mut a = [2,1,3,6,5,7,4];
    let length = a.len()-1;
    quick_sort(&mut a, 0 , length);
    println!("{:?}",a);
}
