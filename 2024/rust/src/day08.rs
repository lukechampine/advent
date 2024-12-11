use std::collections::HashSet;

use crate::utils;

pub fn part1(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let ants: Vec<(utils::Point, u8)> = g.iter().filter(|&(_, c)| c != b'.').collect();
    ants.iter()
        .flat_map(|&a| ants.iter().map(move |&b| (a, b)))
        .filter_map(|(a, b)| (a.1 == b.1 && a != b).then_some((a.0, b.0)))
        .filter_map(|(a, b)| {
            let r = b.add(b.sub(a));
            g.in_bounds(r).then_some(r)
        })
        .collect::<HashSet<_>>()
        .len()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let ants: Vec<(utils::Point, u8)> = g.iter().filter(|&(_, c)| c != b'.').collect();
    let mut seen = HashSet::new();
    ants.iter()
        .flat_map(|&a| ants.iter().map(move |&b| (a, b)))
        .filter_map(|(a, b)| (a.1 == b.1 && a != b).then_some((a.0, b.0)))
        .for_each(|(a, b)| {
            let d = b.sub(a);
            let mut r = b;
            while g.in_bounds(r) {
                seen.insert(r);
                r = r.add(d)
            }
        });
    seen.len().to_string()
}
