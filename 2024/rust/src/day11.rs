use std::collections::HashMap;

use crate::utils;

fn total_stones(s: i64, steps: i64) -> i64 {
    let mut memo: HashMap<(i64, i64), i64> = HashMap::new();

    fn rec(s: i64, steps: i64, memo: &mut HashMap<(i64, i64), i64>) -> i64 {
        if steps == 0 {
            return 1;
        } else if let Some(&n) = memo.get(&(s, steps)) {
            return n;
        }
        let n = if s == 0 {
            rec(1, steps - 1, memo)
        } else if s.to_string().len() % 2 == 0 {
            let s = s.to_string();
            let (l, r) = s.split_at(s.len() / 2);
            rec(l.parse().unwrap(), steps - 1, memo) + rec(r.parse().unwrap(), steps - 1, memo)
        } else {
            rec(s * 2024, steps - 1, memo)
        };
        memo.insert((s, steps), n);
        n
    }
    rec(s, steps, &mut memo)
}

pub fn part1(input: &str) -> String {
    utils::ints(input)
        .map(|s| total_stones(s, 25))
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    utils::ints(input)
        .map(|s| total_stones(s, 75))
        .sum::<i64>()
        .to_string()
}
