fn parse(input: &str) -> (&str, Vec<Vec<&str>>) {
    let (rules, updates) = input.split_once("\n\n").unwrap();
    (
        rules,
        updates
            .lines()
            .map(|line| line.split(',').collect())
            .collect(),
    )
}

fn less(a: &str, b: &str, rules: &str) -> bool {
    rules.contains(&format!("{}|{}", a, b))
}

pub fn part1(input: &str) -> String {
    let (rules, updates) = parse(input);
    updates
        .iter()
        .filter_map(|u| {
            u.is_sorted_by(|a, b| less(a, b, rules))
                .then_some(u[u.len() / 2].parse::<i64>().unwrap())
        })
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let (rules, mut updates) = parse(input);
    updates
        .iter_mut()
        .filter_map(|u| {
            (!u.is_sorted_by(|a, b| less(a, b, rules))).then_some({
                u.sort_unstable_by(|a, b| {
                    less(a, b, rules)
                        .then_some(std::cmp::Ordering::Less)
                        .unwrap_or(std::cmp::Ordering::Greater)
                });
                u[u.len() / 2].parse::<i64>().unwrap()
            })
        })
        .sum::<i64>()
        .to_string()
}
