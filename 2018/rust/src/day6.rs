use scan_fmt::scan_fmt;

pub struct Coord {
    x: isize,
    y: isize,
}

impl Coord {
    fn dist(&self, x: isize, y: isize) -> isize {
        (self.x-x).abs() + (self.y-y).abs()
    }

    fn new(s: &str) -> Coord {
        let (x,y) = scan_fmt!(s, "{}, {}", isize, isize);
        Coord{
            x: x.unwrap(),
            y: y.unwrap(),
        }
    }
}

#[aoc_generator(day6)]
pub fn input_generator(input: &str) -> Vec<Coord> {
    input.lines().map(Coord::new).collect()
}

#[aoc(day6, part2)]
pub fn part2(input: &[Coord]) -> usize {
    (0..1000).map(|x| {
        (0..1000).map(|y| {
            input.iter().map(|c| c.dist(x, y)).sum::<isize>())
        }).filter(|&d| d < 10000)
          .count()
    }).sum()
}
