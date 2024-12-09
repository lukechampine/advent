pub fn ints(s: &str) -> impl Iterator<Item = i64> + '_ {
    s.split(|c: char| !c.is_ascii_digit())
        .filter_map(|s| s.parse().ok())
}

pub fn in_bounds<T>((x, y): (isize, isize), grid: &Vec<Vec<T>>) -> bool {
    x >= 0 && y >= 0 && y < grid.len() as isize && x < grid[y as usize].len() as isize
}
