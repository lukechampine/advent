use crate::utils;

const MAX_X: isize = 101;
const MAX_Y: isize = 103;

struct Robot {
    pos: utils::Point,
    vel: utils::Point,
}

fn mod_like_go(a: isize, b: isize) -> isize {
    ((a % b) + b) % b
}

impl Robot {
    fn from_string(s: &str) -> Self {
        let ints = utils::ints(s).collect::<Vec<_>>();
        Robot {
            pos: (ints[0], ints[1]).into(),
            vel: (ints[2], ints[3]).into(),
        }
    }

    fn update(&mut self) {
        self.pos = self.pos.add(self.vel);
        self.pos.x = mod_like_go(self.pos.x + MAX_X, MAX_X);
        self.pos.y = mod_like_go(self.pos.y + MAX_Y, MAX_Y);
    }
}

pub fn part1(input: &str) -> String {
    let mut robots = input
        .lines()
        .map(|line| Robot::from_string(line))
        .collect::<Vec<_>>();
    for _ in 0..100 {
        for r in &mut robots {
            r.update();
        }
    }
    vec![
        utils::Point::from((0, 0)),
        utils::Point::from((MAX_X / 2 + 1, 0)),
        utils::Point::from((0, MAX_Y / 2 + 1)),
        utils::Point::from((MAX_X / 2 + 1, MAX_Y / 2 + 1)),
    ]
    .iter()
    .map(|min| {
        robots
            .iter()
            .filter(|r| {
                min.x <= r.pos.x
                    && r.pos.x <= min.x + MAX_X / 2 - 1
                    && min.y <= r.pos.y
                    && r.pos.y <= min.y + MAX_Y / 2 - 1
            })
            .count()
    })
    .product::<usize>()
    .to_string()
}

pub fn part2(input: &str) -> String {
    let mut robots = input
        .lines()
        .map(|line| Robot::from_string(line))
        .collect::<Vec<_>>();

    let mut grid = vec![false; (MAX_X * MAX_Y) as usize];
    std::iter::from_fn(move || {
        grid.fill(false);
        for r in &mut robots {
            r.update();
            grid[(r.pos.y * MAX_X + r.pos.x) as usize] = true;
        }
        Some(grid.windows(2).filter(|w| w[0] != w[1]).count())
    })
    .enumerate()
    .take(10000)
    .min_by_key(|x| x.1)
    .unwrap()
    .0
    .to_string()
}
