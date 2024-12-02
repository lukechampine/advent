pub fn ints(s: &str) -> impl Iterator<Item = i64> + '_ {
    s.split_whitespace().map(|n| n.parse().unwrap())
}
