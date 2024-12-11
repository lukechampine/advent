pub fn ints(s: &str) -> impl Iterator<Item = i64> + '_ {
    s.split(|c: char| !c.is_ascii_digit())
        .filter_map(|s| s.parse().ok())
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

    pub fn at(&self, Point { x, y }: Point) -> u8 {
        self.data[(y as usize) * self.width + (x as usize)]
    }

    pub fn in_bounds(&self, Point { x, y }: Point) -> bool {
        0 <= x && (x as usize) < self.width && 0 <= y && (y as usize) < self.height
    }

    pub fn iter(&self) -> impl Iterator<Item = (Point, u8)> + '_ {
        (0..self.height).flat_map(move |y| {
            (0..self.width).map(move |x| ((x, y).into(), self.at((x, y).into())))
        })
    }

    pub fn adj(&self, p: Point) -> impl Iterator<Item = (Point, u8)> + '_ {
        Point::adj(p).filter_map(|p| self.in_bounds(p).then(|| (p, self.at(p))))
    }
}
