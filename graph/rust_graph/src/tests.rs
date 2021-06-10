#[cfg(test)]
mod tests {

    #[test]
    fn test_everything() {
        let mut g = graph_array_impl::graph::new();
        g.insert_vertex("A").unwrap();
        g.insert_vertex("B").unwrap();
        g.insert_vertex("C").unwrap();
        g.insert_vertex("D").unwrap();
        g.insert_connection("A","B",10);
        g.print_all();
    }
}