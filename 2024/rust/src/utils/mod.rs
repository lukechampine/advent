pub fn ints(s: &str) -> impl Iterator<Item = i64> + '_ {
    s.split(|c: char| !c.is_ascii_digit())
        .filter_map(|s| s.parse().ok())
}
