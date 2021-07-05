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

fn binary_search_recur(a: &[i32], mut start: usize, mut end: usize, x: i32) -> Option<usize> {
    // out of bound ?
    if start > end {
        return None;
    }

    let mid = start + (end-start)/2;

    if a[mid] == x {
        return Some(mid);
    } else if x > a[mid] {
        start = mid + 1;
    } else {
        end = mid - 1;
    }

    return binary_search_recur(a,start,end,x);
}

fn binary_search_first_last(
    A: &[i32], 
    x: i32,
    first_or_last: &str,
    ) -> Option<usize> {
    
    let mut start = 0;
    let mut end = A.len()-1;
    
    let mut result = None;

    while start <= end {
        let mid = (start+end)/2;
            
        if A[mid] == x {
            match first_or_last {
                "first" => {
                    result = Some(mid);
                    end = mid - 1;
                },
                "last" => {
                    result = Some(mid);
                    start = mid + 1;
                },
                _ => {return Some(mid);},
            };
        } else if x < A[mid] {
            end = mid - 1;
        }else {
            start = mid + 1;
        }
    }
    return result;
}
/*
fn return_closure(s: &str) -> Box<dyn Fn(usize)> {
  if s == "first" {
    return Box::new(|mid| {
        end = mid - 1;
        result = Some(mid);
    });
  } else if s == "last" {
      return Box::new(|mid| {
        start = mid + 1;
        result = Some(mid);
    });
  } else {
    Box::new(|mid| {
        result = Some(mid);
    })
  }
}*/

fn find_count(a: &[i32], x: i32) {
    let first = binary_search_first_last(a,x,"first").unwrap();
    let last = binary_search_first_last(a,x,"last").unwrap();

    let count = last - first + 1;
    println!("Count for number {} is {}",x,count);

}

fn find_number_of_rotation(a: &[i32])  {
    
    let mut low = 0;
    let mut high = a.len()-1;
    let mut rotation = 0;

    while low <= high {

        // if it was already sorted
        if a[low] <= a[high] {
            rotation = low;
        }

        let mid = (low + high)/2;
        let next = (mid + 1) % a.len();
        let prev = (mid + a.len() -1) % a.len();

        if a[mid] <= a[next] && a[mid] <= a[prev] {
            rotation =  mid;
            break
        }

        // reducing search space
        println!("{} {}",mid,high);
        if a[mid] <= a[high] {

            high = mid - 1;
        } else if a[mid] >= a[low] {
            low = mid + 1
        }

    }
    
    println!("Rotation happened {} times",rotation);

}

// anti clockwise rotation
fn find_element_in_sorted_list(a: &[i32],x: i32) {
   let mut low = 0;
   let mut high = a.len() -1;

   while low <= high {
       let mid = low + (high - low) / 2;
       if a[mid] == x {
           println!("found element {} in index  {}",x,mid);
           return
       }

       // if right part is sorted since if a[mid] is less than a[high]
       //  meaning values at greater index than mid will be greater
       if a[mid] <= a[high] {

            //checking if x lies in the sorted half
            if x > a[mid] && x <= a[high] {
                low = mid + 1;  // search right
            } else {
                high = mid -1; // search right
            }
       } else {
           // if right half not sorted then left half is sorted
           // check if x lies in the left sorted half
           if a[low] <= x && x < a[mid] {
               high = mid -1;
           } else {
               low = mid + 1;
           }
       }
   }
}
fn main() {
    /*let array = [1,2,3,4,5,6,7,8,9];
    let mut b = array.len();
    let i = match binary_search_recur(&array,0,b,9) {
        Some(v) => v,
        None => panic!("value not found"),
    };*/

    /*
    let array = [1,2,3,4,4,5,6,7,8];
    let i = match  binary_search_first_last(&array,4,"last") {
        Some(v) => v,
        None => panic!("value not found"),
    };
    println!("Position is at {}",i);*/
    let array = [6,7,8,9,1,2,3,4,5];
    //find_count(&array,5);
    //find_number_of_rotation(&array);
    find_element_in_sorted_list(&array,3);
}