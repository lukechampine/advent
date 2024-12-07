use crate::utils;

fn parse(input: &str) -> Vec<(i64, Vec<i64>)> {
    input
        .lines()
        .map(|line| {
            let ints: Vec<i64> = utils::ints(line).collect();
            (ints[0], ints[1..].to_vec())
        })
        .collect()
}

fn can_make(r: i64, acc: i64, nums: &[i64]) -> bool {
    if nums.len() == 0 {
        return r == acc;
    }
    can_make(r, acc + nums[0], &nums[1..]) || can_make(r, acc * nums[0], &nums[1..])
}

fn concat_ints(a: i64, b: i64) -> i64 {
    format! {"{}{}", a, b}.parse().unwrap()
}

fn can_make_concat(r: i64, acc: i64, nums: &[i64]) -> bool {
    if nums.len() == 0 {
        return r == acc;
    }
    can_make_concat(r, acc + nums[0], &nums[1..])
        || can_make_concat(r, acc * nums[0], &nums[1..])
        || can_make_concat(r, concat_ints(acc, nums[0]), &nums[1..])
}

pub fn part1(input: &str) -> String {
    parse(input)
        .iter()
        .filter_map(|(r, nums)| can_make(*r, nums[0], &nums[1..]).then_some(r))
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    parse(input)
        .iter()
        .filter_map(|(r, nums)| can_make_concat(*r, nums[0], &nums[1..]).then_some(r))
        .sum::<i64>()
        .to_string()
}
