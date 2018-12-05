fn invert_case(c: char) -> char {
    if c.is_lowercase() {
        c.to_ascii_uppercase()
    } else {
        c.to_ascii_lowercase()
    }
}

fn reduced_len(s: &str) -> usize {
    s.trim().chars().fold(String::new(), |mut r, c| {
        if r.ends_with(invert_case(c)) {
            r.pop();
        } else {
            r.push(c);
        }
        r
    }).len()
}

#[aoc(day5, part1)]
pub fn part1(input: &str) -> usize {
    reduced_len(input)
}

fn remove_pair(input: &str, p: char) -> String {
    input.chars().filter(|&c| c != p && c != invert_case(p)).collect()
}

#[aoc(day5, part2)]
pub fn part2(input: &str) -> usize {
    return "abcdefghijklmnopqrstuvwxyz".chars()
        .map(|c| remove_pair(input, c))
        .map(|s| reduced_len(s.as_str()))
        .min().unwrap()
}
