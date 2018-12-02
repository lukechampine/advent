use std::collections::HashMap;

fn has_rep(s: &str, n: usize) -> bool {
    let mut counts = HashMap::new();
    for c in s.chars() {
        *counts.entry(c).or_insert(0) += 1;
    }
    counts.values().any(|&c| c == n)
}

#[aoc(day2, part1)]
pub fn part1(input: &str) -> usize {
    let twice = input.lines().filter(|s| has_rep(s, 2)).count();
    let thrice = input.lines().filter(|s| has_rep(s, 3)).count();
    twice * thrice
}

fn delete_diff(s1: &str, s2: &str) -> String{
    return s1.chars().zip(s2.chars())
        .filter_map(|x| {
            if x.0 == x.1 {
                Some(x.0)
            } else {
                None
            }
        }).collect();
}

#[aoc(day2, part2)]
pub fn part2(input: &str) -> String {
    for (i, line) in input.lines().enumerate() {
        let target = line.len() - 1;
        for cmp in input.lines().skip(i+1) {
            let rem = delete_diff(line, cmp);
            if rem.len() == target {
                return rem
            }
        }
    }
    unreachable!()
}
