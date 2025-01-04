use std::collections::{HashMap, VecDeque};

use crate::utils;

fn maze_distances(input: &str) -> Vec<(utils::Point, isize)> {
    let g = utils::Grid::from_string(input);
    let start = g.locate(b'E').unwrap();
    let mut dists = HashMap::from([(start, 0)]);
    let mut queue = VecDeque::new();
    queue.push_back(start);
    while let Some(p) = queue.pop_front() {
        for (adj, c) in g.adj(p) {
            if c == b'#' || dists.get(&adj).map_or(false, |&d| d <= dists[&p]) {
                continue;
            }
            dists.insert(adj, dists[&p] + 1);
            queue.push_back(adj);
        }
    }
    dists.into_iter().collect()
}

pub fn part1(input: &str) -> String {
    let dists = maze_distances(input);
    utils::pairs(&dists, &dists)
        .iter()
        .filter(|&&((p, pd), (q, qd))| {
            let d = p.dist(q);
            d <= 2 && pd - (qd + d) >= 100
        })
        .count()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let dists = maze_distances(input);
    utils::pairs(&dists, &dists)
        .iter()
        .filter(|&&((p, pd), (q, qd))| {
            let d = p.dist(q);
            d <= 20 && pd - (qd + d) >= 100
        })
        .count()
        .to_string()
}
