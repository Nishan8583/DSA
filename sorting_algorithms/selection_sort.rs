// section_sort algorithm implementation
fn selection_sort(arr: &mut [i32]) {
    
    // loop through, 0 to last -1 cause the last value will not be needed to sorted
    for i in 0..arr.len()-1 {
        
        // it will hold the currnet first value
        let mut i_min = i;

        // looping current postiion to last
        for j in i..arr.len() {
            if arr[j] < arr[i_min] {
                i_min = j;
            }
        }
        let temp = arr[i];
        arr[i] = arr[i_min];
        arr[i_min] = temp;
    }

}

fn main() {
    let mut arr: [i32;10] = [6,3,8,9,4,5,1,2,6,7];
    selection_sort(&mut arr);
    println!("{:?}",arr);
}