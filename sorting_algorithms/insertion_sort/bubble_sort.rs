fn bubble_sort(a: &mut [i32]) {
    
    // to len()-1 cause we gonna use this to get the position in array
    for i in 0..a.len()-1 {
        // to len()-2 cause we will use this to get the adjacent element, +1 
        for k in 0..a.len()-1 {
            println!("index {}",k);
            if a[k] > a[k+1] {
                let temp = a[k];
                a[k] =  a[k+1];
                a[k+1] = temp;
            }
        }
    }
}

fn main() {
    let mut array = [2,7,4,1,3];
    bubble_sort(&mut array);
    println!("{:?} {}",array,array.len());
    
    // here loop does not go till 4, only till three
    for i in 0..4 {
        println!("yo {}",i);
    }
}