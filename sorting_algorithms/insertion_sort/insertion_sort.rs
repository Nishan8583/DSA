fn insertion_sort(A: &mut [i32]) {
    for i in 1..A.len() {
        let value = A[i];
        let mut hole = i;

        // check if the final value of prvious sorted part is greater,
        // if so, only then we need to sort again
        while (hole >0 && A[hole-1] > value) {
            A[hole] = A[hole-1];  // shifting the greater value to right
            hole = hole -1;
        }
        println!("{} {} {}",hole,A[hole],value);
        A[hole] = value;  // shifting the smaller value to left
    }
}

fn main() {
    let mut r = [7,2,5,3,1];
    insertion_sort(&mut r);
    println!("{:?}",r);
}