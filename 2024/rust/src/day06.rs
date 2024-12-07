use std::collections::HashSet;

fn parse(input: &str) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}

fn start(grid: &Vec<Vec<char>>) -> (isize, isize) {
    for y in 0..grid.len() {
        for x in 0..grid[y].len() {
            if grid[y][x] == '^' {
                return (x as isize, y as isize);
            }
        }
    }
    unreachable!()
}

fn visit(grid: &Vec<Vec<char>>) -> HashSet<(isize, isize)> {
    let (mut x, mut y) = start(&grid);
    let mut dir = (0, -1);
    let mut seen = HashSet::new();
    loop {
        let (nx, ny) = (x + dir.0, y + dir.1);
        if nx < 0 || nx >= grid[0].len() as isize || ny < 0 || ny >= grid.len() as isize {
            return seen;
        }
        if grid[ny as usize][nx as usize] == '#' {
            // turn right
            dir = match dir {
                (0, -1) => (1, 0),
                (1, 0) => (0, 1),
                (0, 1) => (-1, 0),
                (-1, 0) => (0, -1),
                _ => unreachable!(),
            };
        } else {
            (x, y) = (nx, ny);
            seen.insert((x, y));
        }
    }
}

pub fn part1(input: &str) -> String {
    let grid = parse(input);
    visit(&grid).len().to_string()
}

fn has_cycle(grid: &Vec<Vec<char>>, cx: isize, cy: isize) -> bool {
    let (mut x, mut y) = start(&grid);
    let mut dir = (0, -1);
    let mut seen = HashSet::new();
    loop {
        let (nx, ny) = (x + dir.0, y + dir.1);
        if nx < 0 || nx >= grid[0].len() as isize || ny < 0 || ny >= grid.len() as isize {
            return false;
        }
        if grid[ny as usize][nx as usize] == '#' || (nx, ny) == (cx, cy) {
            if !seen.insert((x, y, dir)) {
                return true;
            }
            // turn right
            dir = match dir {
                (0, -1) => (1, 0),
                (1, 0) => (0, 1),
                (0, 1) => (-1, 0),
                (-1, 0) => (0, -1),
                _ => unreachable!(),
            };
        } else {
            (x, y) = (nx, ny);
        }
    }
}

pub fn part2(input: &str) -> String {
    let grid = parse(input);
    let start = start(&grid);
    visit(&grid)
        .iter()
        .filter(|&&(x, y)| (x, y) != start && has_cycle(&grid, x, y))
        .count()
        .to_string()
}
