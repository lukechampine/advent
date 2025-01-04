use std::hash::Hash;
use std::{
    collections::{HashSet, VecDeque},
    ops::{Index, IndexMut},
};

pub fn ints(s: &str) -> impl Iterator<Item = i64> + '_ {
    s.split(|c: char| !c.is_ascii_digit() && c != '-')
        .filter_map(|s| s.parse().ok())
}

pub fn pairs<T: Clone, U: Clone>(a: &Vec<T>, b: &Vec<U>) -> Vec<(T, U)> {
    a.iter()
        .flat_map(|x| b.iter().map(|y| (x.clone(), y.clone())))
        .collect()
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct Point {
    pub x: isize,
    pub y: isize,
}

impl<T> From<(T, T)> for Point
where
    T: TryInto<isize>,
    T::Error: std::fmt::Debug,
{
    fn from((x, y): (T, T)) -> Self {
        Point {
            x: x.try_into().unwrap(),
            y: y.try_into().unwrap(),
        }
    }
}

impl Point {
    pub fn add(self, Point { x, y }: Point) -> Self {
        Point {
            x: self.x + x,
            y: self.y + y,
        }
    }

    pub fn sub(self, Point { x, y }: Point) -> Self {
        Point {
            x: self.x - x,
            y: self.y - y,
        }
    }

    pub fn dist(self, p: Point) -> isize {
        let rel = self.sub(p);
        rel.x.abs() + rel.y.abs()
    }

    pub fn adj(p: Point) -> impl Iterator<Item = Point> {
        [(0, 1), (1, 0), (0, -1), (-1, 0)]
            .iter()
            .map(move |&d| p.add(d.into()))
    }
}

pub struct Grid {
    data: Vec<u8>,
    width: usize,
    height: usize,
}

impl Grid {
    pub fn new(width: usize, height: usize, init: u8) -> Self {
        Grid {
            data: vec![init; width * height],
            width,
            height,
        }
    }

    pub fn from_string(s: &str) -> Self {
        let lines: Vec<&str> = s.lines().collect();
        let height = lines.len();
        let width = lines[0].len();
        let mut data = Vec::with_capacity(width * height);
        for line in lines {
            data.extend(line.bytes());
        }
        Grid {
            data,
            width,
            height,
        }
    }

    pub fn in_bounds(&self, Point { x, y }: Point) -> bool {
        0 <= x && (x as usize) < self.width && 0 <= y && (y as usize) < self.height
    }

    pub fn iter(&self) -> impl Iterator<Item = (Point, u8)> + '_ {
        (0..self.height)
            .flat_map(move |y| (0..self.width).map(move |x| ((x, y).into(), self[(x, y).into()])))
    }

    pub fn adj(&self, p: Point) -> impl Iterator<Item = (Point, u8)> + '_ {
        Point::adj(p).filter_map(|p| self.in_bounds(p).then(|| (p, self[p])))
    }

    pub fn locate(&self, c: u8) -> Option<Point> {
        self.iter().find(|&(_, b)| b == c).map(|(p, _)| p)
    }

    #[allow(dead_code)]
    pub fn render(&self) -> String {
        self.data
            .chunks(self.width)
            .map(|line| String::from_utf8_lossy(line))
            .collect::<Vec<_>>()
            .join("\n")
    }
}

impl Index<Point> for Grid {
    type Output = u8;

    fn index(&self, p: Point) -> &Self::Output {
        &self.data[p.y as usize * self.width + p.x as usize]
    }
}

impl IndexMut<Point> for Grid {
    fn index_mut(&mut self, p: Point) -> &mut Self::Output {
        &mut self.data[p.y as usize * self.width + p.x as usize]
    }
}

pub fn flood<T: Eq + Hash + Copy>(
    start: T,
    mut is_open: impl FnMut(T) -> bool,
    mut next: impl FnMut(T) -> Vec<T>,
) -> impl Iterator<Item = T> {
    let mut seen = HashSet::from([start]);
    let mut queue = VecDeque::new();
    queue.push_back(start);
    std::iter::from_fn(move || {
        let p = queue.pop_front()?;
        for adj in next(p) {
            if is_open(adj) && seen.insert(adj) {
                queue.push_back(adj);
            }
        }
        Some(p)
    })
}
