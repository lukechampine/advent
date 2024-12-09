use std::collections::HashSet;

use crate::utils;

fn parse(input: &str) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}

pub fn part1(input: &str) -> String {
    let grid = parse(input);
    let ants: Vec<(usize, usize)> = grid
        .iter()
        .enumerate()
        .flat_map(move |(y, row)| (0..row.len()).map(move |x| (x, y)))
        .filter(|&(x, y)| grid[y][x] != '.')
        .collect();
    ants.iter()
        .flat_map(|&a| ants.iter().map(move |&b| (a, b)))
        .filter(|&(a, b)| grid[a.1][a.0] == grid[b.1][b.0] && a != b)
        .filter_map(|((ax, ay), (bx, by))| {
            let rx = 2 * bx as isize - ax as isize;
            let ry = 2 * by as isize - ay as isize;
            utils::in_bounds((rx, ry), &grid).then_some((rx, ry))
        })
        .collect::<HashSet<_>>()
        .len()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let grid = parse(input);
    let ants: Vec<(usize, usize)> = grid
        .iter()
        .enumerate()
        .flat_map(move |(y, row)| (0..row.len()).map(move |x| (x, y)))
        .filter(|&(x, y)| grid[y][x] != '.')
        .collect();
    let mut seen = HashSet::new();
    ants.iter()
        .flat_map(|&a| ants.iter().map(move |&b| (a, b)))
        .filter(|&(a, b)| grid[a.1][a.0] == grid[b.1][b.0] && a != b)
        .for_each(|((ax, ay), (bx, by))| {
            let dx = bx as isize - ax as isize;
            let dy = by as isize - ay as isize;
            let (mut rx, mut ry) = (ax as isize + dx, ay as isize + dy);
            while utils::in_bounds((rx, ry), &grid) {
                seen.insert((rx, ry));
                rx += dx;
                ry += dy;
            }
        });
    seen.len().to_string()
}
