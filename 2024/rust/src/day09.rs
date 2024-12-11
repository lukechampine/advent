#[derive(Clone, Copy)]
struct Run {
    id: i64,
    len: i64,
}

fn parse(input: &str) -> Vec<Run> {
    let mut id = 0;
    input
        .trim()
        .chars()
        .enumerate()
        .map(|(i, c)| {
            let n = c.to_digit(10).unwrap() as i64;
            if i % 2 == 0 {
                id += 1;
                Run { id: id - 1, len: n }
            } else {
                Run { id: -1, len: n }
            }
        })
        .collect()
}

pub fn part1(input: &str) -> String {
    let mut fs = parse(input)
        .iter()
        .flat_map(|r| vec![r.id; r.len as usize])
        .collect::<Vec<_>>();
    let (mut i, mut j) = (0, fs.len() - 1);
    while i < j {
        if fs[i] != -1 {
            i += 1;
        } else if fs[j] == -1 {
            j -= 1;
        } else {
            fs.swap(i, j);
            i += 1;
            j -= 1;
        }
    }
    fs.iter()
        .enumerate()
        .filter_map(|(i, &id)| (id != -1).then_some(i as i64 * id))
        .sum::<i64>()
        .to_string()
}

pub fn part2(input: &str) -> String {
    let mut runs = parse(input);

    for j in (1..runs.len()).rev() {
        if runs[j].id == -1 {
            continue;
        }
        if let Some(i) = (0..j).find(|&i| runs[i].id == -1 && runs[i].len >= runs[j].len) {
            let gap = runs[i].len - runs[j].len;
            runs[i] = runs[j];
            runs[j].id = -1;
            runs.insert(i + 1, Run { id: -1, len: gap });
        }
    }

    runs.iter()
        .flat_map(|r| vec![r.id; r.len as usize])
        .enumerate()
        .filter_map(|(i, id)| (id != -1).then_some(i as i64 * id))
        .sum::<i64>()
        .to_string()
}
