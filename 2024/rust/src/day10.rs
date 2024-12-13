use std::collections::HashSet;

use crate::utils;

pub fn part1(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    g.iter()
        .map(|(p, c)| {
            if c != b'0' {
                return 0;
            }
            let mut seen = HashSet::new();
            let mut stack = vec![p];
            let mut sum = 0;
            while let Some(p) = stack.pop() {
                if !seen.insert(p) {
                    continue;
                } else if g[p] == b'9' {
                    sum += 1;
                    continue;
                }
                g.adj(p).for_each(|(adj, d)| {
                    if d == g[p] + 1 {
                        stack.push(adj);
                    }
                });
            }
            sum
        })
        .sum::<usize>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    g.iter()
        .map(|(p, c)| {
            if c != b'0' {
                return 0;
            }
            let mut stack = vec![p];
            let mut sum = 0;
            while let Some(p) = stack.pop() {
                if g[p] == b'9' {
                    sum += 1;
                    continue;
                }
                g.adj(p).for_each(|(adj, d)| {
                    if d == g[p] + 1 {
                        stack.push(adj);
                    }
                });
            }
            sum
        })
        .sum::<usize>()
        .to_string()
}
