use crate::utils;

struct Machine {
    a: utils::Point,
    b: utils::Point,
    prize: utils::Point,
}

impl Machine {
    fn from_string(s: &str) -> Self {
        let ints: Vec<_> = utils::ints(s).collect();
        Machine {
            a: (ints[0], ints[1]).into(),
            b: (ints[2], ints[3]).into(),
            prize: (ints[4], ints[5]).into(),
        }
    }

    fn fracs(&self) -> (isize, isize, isize, isize) {
        let n = self.a.x * self.prize.y - self.a.y * self.prize.x;
        let d = self.a.x * self.b.y - self.a.y * self.b.x;
        ((self.prize.y - self.b.y * n / d), self.a.y, n, d)
    }

    fn is_winnable(&self) -> bool {
        let (n0, d0, n1, d1) = self.fracs();
        n0 % d0 == 0 && n1 % d1 == 0
    }

    fn win_cost(&self) -> isize {
        let (n0, d0, n1, d1) = self.fracs();
        3 * n0 / d0 + n1 / d1
    }
}

pub fn part1(input: &str) -> String {
    input
        .split("\n\n")
        .map(|group| Machine::from_string(group))
        .filter_map(|m| m.is_winnable().then_some(m.win_cost()))
        .sum::<isize>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    input
        .split("\n\n")
        .map(|group| Machine::from_string(group))
        .map(|m| Machine {
            prize: m.prize.add((1e13 as i64, 1e13 as i64).into()),
            ..m
        })
        .filter_map(|m| m.is_winnable().then_some(m.win_cost()))
        .sum::<isize>()
        .to_string()
}
