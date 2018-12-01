use std::collections::HashSet;

#[aoc(day1, part1)]
pub fn part1(input: &str) -> isize {
    input.lines().map(|x| x.parse::<isize>().unwrap()).sum()
}

#[aoc(day1, part2)]
pub fn part2(input: &str) -> isize {
    let mut seen = HashSet::new();
    return input
        .lines()
        .map(|x| x.parse::<isize>().unwrap())
        .cycle()
        .scan(0, |state, n| { *state += n; Some(*state) })
        .find(|n| !seen.insert(*n)).unwrap();
}
