use crate::utils;
use std::collections::HashMap;

fn parse(input: &str) -> Vec<utils::Point> {
    utils::ints(input)
        .collect::<Vec<i64>>()
        .chunks(2)
        .map(|w| (w[0], w[1]).into())
        .collect()
}

fn best_path(bytes: Vec<utils::Point>) -> Option<i64> {
    let mut grid = utils::Grid::new(71, 71, b'.');
    bytes.iter().for_each(|&p| grid[p] = b'#');

    let mut best: Option<i64> = None;
    let mut dists = HashMap::new();
    let mut queue = vec![((0, 0).into(), 0)];
    while let Some((p, d)) = queue.pop() {
        if p == (70, 70).into() {
            best = Some(best.map_or(d, |b| b.min(d)));
            continue;
        }
        grid.adj(p).for_each(|(m, c)| {
            let e = dists.entry(m).or_insert(i64::MAX);
            if c != b'#' && d + 1 < *e {
                *e = d + 1;
                queue.push((m, d + 1));
            }
        });
    }
    best
}

pub fn part1(input: &str) -> String {
    best_path(parse(input)[..1024].into()).unwrap().to_string()
}

pub fn part2(input: &str) -> String {
    let bytes = parse(input);
    let n = (0..bytes.len())
        .collect::<Vec<_>>()
        .partition_point(|&n| best_path(bytes[0..n].into()).is_some());
    format!("{},{}", bytes[n - 1].x, bytes[n - 1].y)
}
