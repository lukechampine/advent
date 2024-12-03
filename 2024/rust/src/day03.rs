use crate::utils;
use regex::Regex;

fn parse(input: &str) -> Vec<&str> {
    let re = Regex::new(r"mul\((\d+),(\d+)\)|do\(\)|don't\(\)").unwrap();
    re.find_iter(input).map(|m| m.as_str()).collect()
}

pub fn part1(input: &str) -> String {
    parse(input)
        .iter()
        .map(|m| {
            if m.starts_with("mul") {
                utils::ints(m).product()
            } else {
                0
            }
        })
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    parse(input)
        .iter()
        .fold((0, true), |(acc, enabled), m| {
            if m.starts_with("do") {
                (acc, m.starts_with("do()"))
            } else {
                (
                    acc + enabled as i64 * utils::ints(m).product::<i64>(),
                    enabled,
                )
            }
        })
        .0
        .to_string()
}
