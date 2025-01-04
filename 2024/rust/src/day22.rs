use std::collections::{HashMap, HashSet};

use crate::utils;

fn next(mut x: i64) -> i64 {
    x ^= x << 6;
    x &= (1 << 24) - 1;
    x ^= x >> 5;
    x &= (1 << 24) - 1;
    x ^= x << 11;
    x &= (1 << 24) - 1;
    x
}

pub fn part1(input: &str) -> String {
    let secrets = utils::ints(input);
    secrets
        .map(|mut n| {
            for _ in 0..2000 {
                n = next(n)
            }
            return n;
        })
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let secrets = utils::ints(input);
    let mut sales = HashMap::new();
    for mut n in secrets {
        let mut prev_price = 0;
        let mut deltas = [0; 4];
        let mut seen = HashSet::new();
        for i in 0..=2000 {
            let price = n % 10;
            n = next(n);
            deltas.rotate_left(1);
            deltas[3] = price - prev_price;
            prev_price = price;
            if i >= 4 && !seen.contains(&deltas) {
                *sales.entry(deltas).or_insert(0) += price;
                seen.insert(deltas);
            }
        }
    }
    sales.values().max().unwrap().to_string()
}
