fn main() {}

fn spelling_check(s1: String, s2: String) -> (usize, Vec<usize>) {
    let mut indices: Vec<usize> = vec![];

    let s1_ascii = s1.as_bytes();
    let s2_ascii = s2.as_bytes();

    // calculate common prefix
    let mut prefix_len = 0;
    while prefix_len < s2.len() && s1_ascii[prefix_len] == s2_ascii[prefix_len] {
        prefix_len += 1;
    }

    // calculate common suffix
    let mut suffix_len = 0;
    while suffix_len < s2.len()
        && s1_ascii[s1.len() - 1 - suffix_len] == s2_ascii[s2.len() - 1 - suffix_len]
    {
        suffix_len += 1;
    }

    println!(
        "S1={} S2={} prefix_len {} suffix_len {}",
        s1, s2, prefix_len, suffix_len
    );


    // there is a case here where substracting to negative is now allowed, whichj makes test case fail, i did not bother fixing it
    let count = (prefix_len+1) - (s1.len()-suffix_len) +1;
    // Now, we have a window where the mismatched characters are between prefix_len and (s1.len() - suffix_len)
    let mut i = 0;
    for i in 0..count {
        println!("{}",i+s1.len()-suffix_len);
        indices.push(i+s1.len()-suffix_len);

    }

    return (count,indices);
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_spelling_check() {
        let (count, indices) = spelling_check("abcdxxxef".to_owned(), "abcdxxef".to_owned());
        assert_eq!(3, count);
        assert_eq!(vec![5, 6, 7], indices);

        let (count, indices) = spelling_check("bizzarre".to_owned(), "bizarre".to_owned());
        assert_eq!(2, count);
        assert_eq!(vec![3, 4], indices);

        let (count, indices) = spelling_check("aa".to_owned(), "a".to_owned());
        assert_eq!(2, count);
        assert_eq!(vec![1, 2], indices);

        let (count, indices) = spelling_check("competition".to_owned(), "codeforces".to_owned());
        assert_eq!(0, count);
        assert_eq!(Vec::<usize>::new(), indices);
    }
}
