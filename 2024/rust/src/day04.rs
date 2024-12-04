fn stamps(grid: Vec<&str>, stamp: Vec<(usize, usize)>) -> Vec<String> {
    grid.iter()
        .enumerate()
        .flat_map(move |(y, row)| (0..row.len()).map(move |x| (x, y)))
        .filter_map(|(x, y)| {
            let stamped: String = stamp
                .iter()
                .filter_map(|(dx, dy)| grid.get(y + dy).and_then(|row| row.chars().nth(x + dx)))
                .collect();
            (stamped.len() == stamp.len()).then_some(stamped)
        })
        .collect()
}

pub fn part1(input: &str) -> String {
    [
        [(0, 0), (1, 0), (2, 0), (3, 0)],
        [(0, 0), (0, 1), (0, 2), (0, 3)],
        [(0, 0), (1, 1), (2, 2), (3, 3)],
        [(0, 3), (1, 2), (2, 1), (3, 0)],
    ]
    .iter()
    .map(|stamp| {
        stamps(input.lines().collect(), stamp.to_vec())
            .iter()
            .filter(|&s| s == "XMAS" || s == "SAMX")
            .count() as i64
    })
    .sum::<i64>()
    .to_string()
}

pub fn part2(input: &str) -> String {
    let cross = vec![(0, 0), (2, 0), (1, 1), (0, 2), (2, 2)];
    stamps(input.lines().collect(), cross)
        .iter()
        .filter(|&s| s == "MSAMS" || s == "SMASM" || s == "MMASS" || s == "SSAMM")
        .count()
        .to_string()
}
