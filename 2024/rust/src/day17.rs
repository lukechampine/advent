use crate::utils;

fn run(regs: Vec<i64>, prog: Vec<i64>) -> Vec<i64> {
    let mut a = regs[0];
    let mut b = regs[1];
    let mut c = regs[2];

    let mut out = Vec::new();
    let mut i = 0;
    while i < prog.len() {
        let arg = prog[i + 1];
        let combo = match arg {
            (0..=3) => arg,
            4 => a,
            5 => b,
            6 => c,
            _ => unreachable!(),
        } % 8;
        match prog[i] {
            0 => a /= 1 << combo,
            1 => b ^= arg,
            2 => b = combo,
            3 => i = if a == 0 { i } else { arg as usize - 2 },
            4 => b ^= c,
            5 => out.push(combo),
            6 => b = a / (1 << combo),
            7 => c = a / (1 << combo),
            _ => unreachable!(),
        }
        i += 2;
    }
    out
}

fn parse(input: &str) -> (Vec<i64>, Vec<i64>) {
    let ints = utils::ints(input).collect::<Vec<i64>>();
    (ints[..3].to_vec(), ints[3..].to_vec())
}

pub fn part1(input: &str) -> String {
    let (regs, prog) = parse(input);
    let out = run(regs, prog);
    out.iter()
        .map(|n| n.to_string())
        .collect::<Vec<String>>()
        .join(",")
}

pub fn part2(_input: &str) -> String {
    "TODO".to_string()
}
