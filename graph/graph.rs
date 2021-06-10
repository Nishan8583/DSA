mod graph{
    pub struct GraphArray{
        vertex_list: Vec<String>,
        edge_list: Vec<Edges>,
    }

    pub struct Edges{
        start_edge: String,
        end_edge: String,
        cost: String,
    }

    pub fn new() -> GraphArray {
        GraphArray{
            vertex_list: vec![],
            edge_list: vec![],
        }
    }

    impl GraphArray{
        pub fn Insert(&mut self, vertex_name: String) 
    }
}

fn main() {
    let mut g = graph::new();

}