// graph_matrix module implements adjancy matrix graph

// GMatrix is the structure that will implement the matrix
pub struct GMatrix{
    pub vec_list: Vec<String>,
    pub matrix: Vec<Vec<i64>>,
}

// new returns a GMatrix structure
pub fn new() -> GMatrix {
    let v: Vec<Vec<i64>> = Vec::new();
    GMatrix{
        vec_list: vec![],
        matrix: v,
    }
}


// methods for GMatrix structure
impl GMatrix{

    // insert_vertex inserts into the vertex
    pub fn insert_vertex(&mut self,vertex_name: &str) -> Result<(),String> {
       
        // check if vertex already present
        if self.vec_list.contains(&vertex_name.to_string()){
            return Err(format!("Vector already present"));
        }

        // pushing the new vertex to vector
        self.vec_list.push(vertex_name.to_string());

        // creating a new vector 
        self.matrix.push(vec![-1;self.vec_list.len()]);
        self.update_vector();
        Ok(())
    }

    /// update_vector adds another row when another vector is added
    /// will be called inside insert_vector function, so no need to 
    /// be public
    fn update_vector(&mut self) {
        for i in 0..self.vec_list.len(){
            if self.matrix[i].len() < self.vec_list.len() {
                self.matrix[i].push(-1);
            }
        }
    }

    // print_connections prints all connections
    pub fn print_connections(&self) {
        for (pos,v) in self.vec_list.iter().enumerate() {
            println!("Connection for {} are {:?}",v,self.matrix[pos]);
        }
    }

    // add_connections updates the connection i.e. A[i][j] position to some other value than -1
    // with the actual costs
    pub fn add_connections(&mut self, source_vertex: &str, destination_vertex: &str, cost: i64) -> Result<(),String> {
        let fvalue = self.vec_list.iter().position(|r| r == source_vertex).unwrap();
        let lvalue = self.vec_list.iter().position(|r| r == destination_vertex).unwrap();

        self.matrix[fvalue][lvalue] = cost;
        Ok(())
    }

    // get_cost returns the cost of connection
    pub fn get_cost(&self, source_vertex: &str, destination_vertex: &str) -> i64 {
        let fvalue = self.vec_list.iter().position(|r| r == source_vertex).unwrap();
        let lvalue = self.vec_list.iter().position(|r| r == destination_vertex).unwrap();
        self.matrix[fvalue][lvalue]
    }


}
