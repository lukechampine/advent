use std::fs;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod utils;

fn main() {
    let day = std::env::args().nth(1).unwrap().parse::<u32>().unwrap();
    let input = fs::read_to_string(format!("../day{}_input.txt", day)).unwrap();
    let (part1, part2) = match day {
        1 => (day01::part1(&input), day01::part2(&input)),
        2 => (day02::part1(&input), day02::part2(&input)),
        3 => (day03::part1(&input), day03::part2(&input)),
        4 => (day04::part1(&input), day04::part2(&input)),
        5 => (day05::part1(&input), day05::part2(&input)),
        _ => unimplemented!(),
    };
    println!("{}\n{}", part1, part2);
}
