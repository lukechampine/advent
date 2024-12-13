use std::collections::HashSet;

use crate::utils;

pub fn part1(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let mut seen = HashSet::new();
    g.iter()
        .map(|(p, c)| {
            if !seen.insert(p) {
                return 0;
            }
            let mut area = 0;
            let mut perimeter = 0;
            utils::flood(
                p,
                |p| g.in_bounds(p) && g[p] == c,
                |p| {
                    g.adj(p)
                        .filter_map(|(a, _)| (g[a] == c).then_some(a))
                        .collect()
                },
            )
            .for_each(|p| {
                seen.insert(p);
                area += 1;
                perimeter += 4 - g.adj(p).filter(|&(a, _)| (g[a] == c)).count();
            });
            area * perimeter
        })
        .sum::<usize>()
        .to_string()
}

fn count_sides(shape: Vec<utils::Point>) -> usize {
    let mut sides = 0;
    let mut seen = HashSet::new();
    for d in vec![
        ((-1, 0).into(), (0, 1).into(), (0, -1).into()),
        ((1, 0).into(), (0, 1).into(), (0, -1).into()),
        ((0, -1).into(), (1, 0).into(), (-1, 0).into()),
        ((0, 1).into(), (1, 0).into(), (-1, 0).into()),
    ] {
        for &p in &shape {
            if seen.contains(&p) || shape.contains(&p.add(d.0)) {
                continue;
            }
            sides += 1;
            seen.extend(utils::flood(
                p,
                |p| shape.contains(&p) && !shape.contains(&p.add(d.0)),
                |p| vec![p.add(d.1), p.add(d.2)],
            ));
        }
        seen.clear();
    }
    sides
}

pub fn part2(input: &str) -> String {
    let g = utils::Grid::from_string(input);
    let mut seen = HashSet::new();
    g.iter()
        .map(|(p, c)| {
            if !seen.insert(p) {
                return 0;
            }
            let shape: Vec<_> = utils::flood(
                p,
                |p| g.in_bounds(p) && g[p] == c,
                |p| {
                    g.adj(p)
                        .filter_map(|(a, _)| (g[a] == c).then_some(a))
                        .collect()
                },
            )
            .collect();
            shape.iter().for_each(|&p| {
                seen.insert(p);
            });

            shape.len() * count_sides(shape)
        })
        .sum::<usize>()
        .to_string()
}
