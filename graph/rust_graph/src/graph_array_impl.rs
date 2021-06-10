pub mod graph{
    pub struct GraphArray{
        vertex_list: Vec<String>,
        edge_list: Vec<Edges>,
    }

    pub struct Edges{
        start_edge: String,
        end_edge: String,
        cost: i64,
    }

    pub fn new() -> GraphArray {
        GraphArray{
            vertex_list: vec![],
            edge_list: vec![],
        }
    }

    impl GraphArray{

        // insert_vertex takes a vertex and inserts that into a graph
        pub fn insert_vertex(&mut self, vertex_name: &str) -> Result<(),String> {
            for (i,v) in self.vertex_list.iter().enumerate(){
                if *v == vertex_name{
                    return Err(format!("Key {:?} already present",vertex_name));
                }
            }
            self.vertex_list.push(vertex_name.to_string());
            return Ok(());
        }

        // insert_connection takes two vertices and a costs, and adds a connection
        // between them with the given associated cost
        pub fn insert_connection(&mut self, 
            start_edge: &str, 
            end_edge: &str,
            cost: i64) {
                self.edge_list.push(
                    Edges{
                        start_edge: start_edge.to_string(),
                        end_edge:end_edge.to_string(),
                        cost,
                    }
                );
            }
        
        // pritn_all prints all vertices and edges
        pub fn print_all(&self){
            for (pos,v) in self.vertex_list.iter().enumerate(){
                println!("The vertex {} in position {}",v,pos);
            }

            for (pos,v) in self.edge_list.iter().enumerate(){
                println!("Connections from {} to {} costs {}",v.start_edge,
            v.end_edge,v.cost);
            }
        }

        // check_if_connected takes two vertices, if connected returns its cost
        // else returns -1
        pub fn check_if_connected(&self, start_vertex: String, end_vertex: String) -> i64{
            for (pos,v) in self.edge_list.iter().enumerate() {
                if v.start_edge == start_vertex && 
                v.end_edge == end_vertex{
                    return v.cost;
                }
            }
            return -1;
        }

        // find_shortest_path takes two vertices and returns the shortest path
        // between them, lmnao  useless function
        /*
        pub fn find_shortest_path(&self, start_vertex: String,
        end_vertex: String) -> i64{

            let mut shortest = -1;

            for (_,v) in self.edge_list.iter().enumerate() {
                if v.start_edge == start_vertex && v.end_edge == end_vertex {
                    if v.cost < shortest {
                        shortest = v.cost;
                    }
                }
            }

            shortest
        }*/
    }   
}


