use std::collections::HashMap;

fn possible(design: &str, patterns: &[&str]) -> bool {
    patterns.iter().any(|&p| {
        let rem = design.strip_prefix(p).unwrap_or(design);
        rem.len() < design.len() && (design == p || possible(rem, patterns))
    })
}

fn possibilities(design: &str, patterns: &[&str], memo: &mut HashMap<String, usize>) -> usize {
    if let Some(&count) = memo.get(design) {
        count
    } else {
        let count = patterns
            .iter()
            .map(|p| match design.strip_prefix(p) {
                None => 0,
                Some("") => 1,
                Some(remaining) => possibilities(remaining, patterns, memo),
            })
            .sum();
        memo.insert(design.to_string(), count);
        count
    }
}

fn parse(input: &str) -> (Vec<&str>, Vec<&str>) {
    let groups = input.split("\n\n").collect::<Vec<&str>>();
    (groups[0].split(", ").collect(), groups[1].lines().collect())
}

pub fn part1(input: &str) -> String {
    let (patterns, designs) = parse(input);
    designs
        .iter()
        .filter(|d| possible(d, &patterns))
        .count()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let (patterns, designs) = parse(input);
    designs
        .iter()
        .map(|d| possibilities(d, &patterns, &mut HashMap::new()))
        .sum::<usize>()
        .to_string()
}
