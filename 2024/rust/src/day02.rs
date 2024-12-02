use crate::utils;

fn parse(input: &str) -> Vec<Vec<i64>> {
    input
        .lines()
        .map(|line| utils::ints(line).collect())
        .collect()
}

fn safe(report: &[i64]) -> bool {
    let deltas: Vec<i64> = report.windows(2).map(|w| w[1] - w[0]).collect();
    (deltas.iter().all(|&d| d > 0) || deltas.iter().all(|&d| d < 0))
        && deltas.iter().map(|d| d.abs()).all(|d| 0 < d && d < 4)
}

pub fn part1(input: &str) -> String {
    let reports = parse(input);
    reports.iter().filter(|&r| safe(r)).count().to_string()
}

pub fn part2(input: &str) -> String {
    let reports = parse(input);
    reports
        .iter()
        .filter(|&r| {
            safe(r)
                || (0..r.len()).any(|i| {
                    let mut r = r.clone();
                    r.remove(i);
                    safe(&r)
                })
        })
        .count()
        .to_string()
}
