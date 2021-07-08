#[derive(Debug)]
enum Address<'a> {
    Ptr(&'a Node<'a>),
    None,
}

#[derive(Debug)]
struct Node<'a> {
    value: i32,
    ptr: Address<'a>,
}


fn reverse_lined_list<'a>(l: &'a mut Node<'a>) {
    let mut next = Address::None;  // will hold the next value
    let mut prev =  &Node{value:0,ptr:Address::None};  // will hold the prev value

    // this will be the current node use in loop
    let mut n = l;

    loop {

        // storing the next pointer
        next = &n.ptr;

        // overwriting the next pointer
        n.ptr = Address::Ptr(prev);

        // this will be the previous value for next loop
        prev = n;

        match next {
            Address::None => {
                l = n;
                return
            },
            Address::Ptr(&ref mut v) => {
                n = v;
            },
        }
    }
}
fn main() {
    let mut n1 = Node{
        value: 32,
        ptr: Address::None,
    };

    let mut head = &mut n1;

    head.ptr = Address::Ptr(&mut Node{
        value: 33,
        ptr: Address::Ptr(
            &mut Node{
                value: 34,
                ptr: Address::None,
            }
        ),
    });

    println!("{:?}",head);
}