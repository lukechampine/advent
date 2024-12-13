use std::collections::HashSet;

use crate::utils;

fn turn_right(dir: utils::Point) -> utils::Point {
    match dir {
        utils::Point { x: 0, y: -1 } => utils::Point { x: 1, y: 0 },
        utils::Point { x: 1, y: 0 } => utils::Point { x: 0, y: 1 },
        utils::Point { x: 0, y: 1 } => utils::Point { x: -1, y: 0 },
        utils::Point { x: -1, y: 0 } => utils::Point { x: 0, y: -1 },
        _ => unreachable!(),
    }
}

fn visit(g: &utils::Grid) -> HashSet<utils::Point> {
    let mut p = g.iter().find(|&(_, c)| c == b'^').unwrap().0;
    let mut dir = utils::Point { x: 0, y: -1 };
    let mut seen = HashSet::new();
    loop {
        let n = p.add(dir);
        if !g.in_bounds(n) {
            return seen;
        } else if g[n] == b'#' {
            dir = turn_right(dir);
        } else {
            p = n;
            seen.insert(p);
        }
    }
}

pub fn part1(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    visit(&g).len().to_string()
}

fn has_cycle(g: &utils::Grid, c: utils::Point) -> bool {
    let mut p = g.iter().find(|&(_, c)| c == b'^').unwrap().0;
    let mut dir = utils::Point { x: 0, y: -1 };
    let mut seen = HashSet::new();
    loop {
        let n = p.add(dir);
        if !g.in_bounds(n) {
            return false;
        } else if g[n] == b'#' || n == c {
            if !seen.insert((p, dir)) {
                return true;
            }
            dir = turn_right(dir);
        } else {
            p = n;
        }
    }
}

pub fn part2(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let start = g.iter().find(|&(_, c)| c == b'^').unwrap().0;
    visit(&g)
        .iter()
        .filter(|&&p| p != start && has_cycle(&g, p))
        .count()
        .to_string()
}
