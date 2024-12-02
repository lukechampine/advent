use std::fs;

mod day01;
mod day02;
mod utils;

fn main() {
    let day = std::env::args().nth(1).unwrap().parse::<u32>().unwrap();
    let input = fs::read_to_string(format!("../day{}_input.txt", day)).unwrap();
    let (part1, part2) = match day {
        1 => (day01::part1(&input), day01::part2(&input)),
        2 => (day02::part1(&input), day02::part2(&input)),
        _ => unimplemented!(),
    };
    println!("{}\n{}", part1, part2);
}
