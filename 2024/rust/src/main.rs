use std::fs;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day06;
mod day07;
mod day08;
mod day09;
mod day10;
mod day11;
mod day12;
mod day13;
mod day14;
mod day17;
mod day18;
mod day19;
mod day20;
mod day22;
mod day25;
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
        6 => (day06::part1(&input), day06::part2(&input)),
        7 => (day07::part1(&input), day07::part2(&input)),
        8 => (day08::part1(&input), day08::part2(&input)),
        9 => (day09::part1(&input), day09::part2(&input)),
        10 => (day10::part1(&input), day10::part2(&input)),
        11 => (day11::part1(&input), day11::part2(&input)),
        12 => (day12::part1(&input), day12::part2(&input)),
        13 => (day13::part1(&input), day13::part2(&input)),
        14 => (day14::part1(&input), day14::part2(&input)),
        17 => (day17::part1(&input), day17::part2(&input)),
        18 => (day18::part1(&input), day18::part2(&input)),
        19 => (day19::part1(&input), day19::part2(&input)),
        20 => (day20::part1(&input), day20::part2(&input)),
        22 => (day22::part1(&input), day22::part2(&input)),
        25 => (day25::part1(&input), day25::part2(&input)),
        _ => unimplemented!(),
    };
    println!("{}\n{}", part1, part2);
}
