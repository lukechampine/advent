use crate::utils;

pub fn part1(input: &str) -> String {
    let mut locks: Vec<[usize; 5]> = Vec::new();
    let mut keys: Vec<[usize; 5]> = Vec::new();
    input.split("\n\n").for_each(|g| {
        let a = std::array::from_fn(|i| {
            g.lines()
                .filter(|line| line.bytes().nth(i) == Some(b'#'))
                .count()
        });
        if g.starts_with("#####") {
            locks.push(a);
        } else {
            keys.push(a);
        }
    });
    utils::pairs(&locks, &keys)
        .iter()
        .filter(|(l, k)| (0..5).all(|i| l[i] + k[i] <= 7))
        .count()
        .to_string()
}

pub fn part2(_input: &str) -> String {
    "Merry Christmas!".to_string()
}
