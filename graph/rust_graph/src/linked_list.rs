#[derive(Clone,Debug)]
pub enum Address {
    Ptr(Box<Node>),
    None,
}

#[derive(Clone,Debug)]
pub struct Node {
    connection: String,
    value: i32,
    next: Address,
}

impl Node {

    pub fn new(connection:String, value: i32) -> Node{
        Node{
            connection,
            value,
            next: Address::None,
        }
    }

    pub fn insert(&mut self, connection: String, value: i32) {
        match self.next {
            Address::Ptr(ref mut v) => {
                v.insert(connection,value);
            }
            Address::None => {
                self.next = Address::Ptr(Box::new(Node{connection,value,next:Address::None}))
            }
        }
    }

    pub fn delete(&mut self, value:i32){

        match self.next{
            Address::Ptr(ref mut v) => {
                if v.value == value {

                    // get the ptr to next address, and put it before
                    // Ex: let a = self.next and b = self.next.next
                    // a = b
                    // now even though b points at self.next.next, which may be c
                    // no one points at b itself, self.next, and thus rust seems to drop 
                    // it knowing that, wow 
                    self.next = v.next.clone();
                } else {
                    v.delete(value)
                }
            }
            Address::None => {
                if self.value == value {
                    self.value = 0
                } else {
                    println!("Elemement does not exist")
                }
            }
        }
    }
}

