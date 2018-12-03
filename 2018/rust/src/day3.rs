use scan_fmt::scan_fmt;

pub struct Claim {
    id: usize,
    x: usize,
    y: usize,
    w: usize,
    h: usize,
}

impl Claim {
    fn apply(&self, grid: &mut[[usize; 1000]; 1000]) {
        for x in 0..self.w {
            for y in 0..self.h {
                grid[self.x+x][self.y+y] += 1
            }
        }
    }

    fn overlaps(&self, grid: &[[usize; 1000]; 1000]) -> bool {
        for x in 0..self.w {
            for y in 0..self.h {
                if grid[self.x+x][self.y+y] > 1 {
                    return true
                }
            }
        }
        return false
    }

    fn new(s: &str) -> Claim {
        let (id,x,y,w,h) = scan_fmt!(s, "#{} @ {},{}: {}x{}", usize, usize, usize, usize, usize);
        Claim{
            id: id.unwrap(),
            x: x.unwrap(),
            y: y.unwrap(),
            w: w.unwrap(),
            h: h.unwrap(),
        }
    }
}

#[aoc_generator(day3)]
pub fn input_generator(input: &str) -> Vec<Claim> {
    input.lines().map(Claim::new).collect()
}

#[aoc(day3, part1)]
pub fn part1(input: &[Claim]) -> usize {
    let mut grid = [[0usize; 1000]; 1000];
    input.iter().for_each(|c| c.apply(&mut grid));
    grid.iter().flat_map(|r| r.iter()).filter(|&&c| c > 1).count()
}

#[aoc(day3, part2)]
pub fn part2(input: &[Claim]) -> usize {
    let mut grid = [[0usize; 1000]; 1000];
    input.iter().for_each(|c| c.apply(&mut grid));
    input.iter().find(|c| !c.overlaps(&grid)).unwrap().id
}
