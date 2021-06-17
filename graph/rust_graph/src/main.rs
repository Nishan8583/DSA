use std::collections::HashMap;

// linked_list to create adjacency list
mod linked_list;

// Graph Adjacency list structure
struct GraphAdj {
    vec_edge: HashMap<String,linked_list::Node>
}

impl GraphAdj {
    pub fn new() -> GraphAdj{
        GraphAdj{
            vec_edge: HashMap::new(),
        }
    }

    pub fn insert(&mut self, value: String) -> Result<(),String> {

        if self.vec_edge.contains_key(&value) {
            return Err(format!("Key already present"))
        }

        self.vec_edge.insert(value,
            linked_list::Node::new(String::from(""),-1)
        );
        Ok(())
    } 

    pub fn add_connection(&mut self, 
        source_vertex: String,
        destination_vertex: String,
        cost: i32,
    ){

        for (key,value) in self.vec_edge.iter_mut() {
            if *key == source_vertex {
                value.insert(destination_vertex.clone(),cost);
            } 
        }
    }

}
fn main(){}