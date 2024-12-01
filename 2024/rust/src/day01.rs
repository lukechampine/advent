fn parse(input: &str) -> (Vec<i64>, Vec<i64>) {
    let (mut left, mut right): (Vec<i64>, Vec<i64>) = input
        .lines()
        .map(|line| {
            let (left, right) = line.split_once("   ").unwrap();
            (left.parse::<i64>().unwrap(), right.parse::<i64>().unwrap())
        })
        .unzip();
    left.sort();
    right.sort();
    (left, right)
}

pub fn part1(input: &str) -> String {
    let (left, right) = parse(input);
    left.iter()
        .zip(right.iter())
        .map(|(&l, &r)| (l - r).abs())
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let (left, right) = parse(input);
    left.iter()
        .map(|l| l * right.iter().filter(|&r| r == l).count() as i64)
        .sum::<i64>()
        .to_string()
}
