use crate::utils;

fn stamps(g: &utils::Grid, stamp: Vec<utils::Point>) -> Vec<String> {
    g.iter()
        .map(|(p, _)| {
            stamp
                .iter()
                .map(|&s| p.add(s))
                .filter_map(|p| g.in_bounds(p).then(|| g[p] as char))
                .collect()
        })
        .collect()
}

pub fn part1(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    [
        [(0, 0), (1, 0), (2, 0), (3, 0)],
        [(0, 0), (0, 1), (0, 2), (0, 3)],
        [(0, 0), (1, 1), (2, 2), (3, 3)],
        [(0, 3), (1, 2), (2, 1), (3, 0)],
    ]
    .iter()
    .map(|stamp| {
        stamps(&g, stamp.map(|p| p.into()).to_vec())
            .iter()
            .filter(|&s| s == "XMAS" || s == "SAMX")
            .count() as i64
    })
    .sum::<i64>()
    .to_string()
}

pub fn part2(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let cross = [(0, 0), (2, 0), (1, 1), (0, 2), (2, 2)];
    stamps(&g, cross.map(|p| p.into()).to_vec())
        .iter()
        .filter(|&s| s == "MSAMS" || s == "SMASM" || s == "MMASS" || s == "SSAMM")
        .count()
        .to_string()
}
