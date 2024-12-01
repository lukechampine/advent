use std::fs;

mod day01;

fn main() {
    let day = std::env::args().nth(1).unwrap().parse::<u32>().unwrap();
    let input = fs::read_to_string(format!("inputs/day{:02}.txt", day)).unwrap();
    let (part1, part2) = match day {
        1 => (day01::part1(&input), day01::part2(&input)),
        _ => unimplemented!(),
    };
    println!("{}\n{}", part1, part2);
}
