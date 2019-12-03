const std = @import("std");

fn readFileLines() ![][]const u8 {
    const data = try std.io.readFileAlloc(std.heap.direct_allocator, "day3_input.txt");
    var it = std.mem.separate(data, "\n");
    var num_lines: usize = 0;
    while (it.next()) |line| {
        if (line.len > 0)
            num_lines += 1;
    }
    var lines = try std.heap.direct_allocator.alloc([]const u8, num_lines);
    it = std.mem.separate(data, "\n");
    num_lines = 0;
    while (it.next()) |line| {
        if (line.len > 0) {
            lines[num_lines] = line;
            num_lines += 1;
        }
    }
    return lines;
}

fn parseInt(comptime T: type, str: []const u8) T {
    return std.fmt.parseInt(T, str, 10) catch unreachable;
}

const Pos = struct {
    x: i32 = 0,
    y: i32 = 0,

    fn move(p: Pos, dir: u8) Pos {
        return switch (dir) {
            'U' => Pos{ .x = p.x, .y = p.y + 1 },
            'D' => Pos{ .x = p.x, .y = p.y - 1 },
            'L' => Pos{ .x = p.x - 1, .y = p.y },
            'R' => Pos{ .x = p.x + 1, .y = p.y },
            else => unreachable,
        };
    }

    fn manhattan_dist(p: Pos, o: Pos) i32 {
        return (std.math.absInt(p.x - o.x) catch unreachable) +
            (std.math.absInt(p.y - o.y) catch unreachable);
    }
};

pub fn main() !void {
    const lines = try readFileLines();
    var map = std.AutoHashMap(Pos, usize).init(std.heap.direct_allocator);
    var p = Pos{};
    var steps: usize = 0;
    var it = std.mem.separate(lines[0], ",");
    while (it.next()) |d| {
        var n = parseInt(i32, d[1..]);
        while (n > 0) : (n -= 1) {
            p = p.move(d[0]);
            steps += 1;
            _ = try map.put(p, steps);
        }
    }

    p = Pos{};
    steps = 0;
    it = std.mem.separate(lines[1], ",");
    var bestDist: i32 = 99999999;
    var bestSteps: usize = 99999999;
    while (it.next()) |d| {
        var n = parseInt(i32, d[1..]);
        while (n > 0) : (n -= 1) {
            p = p.move(d[0]);
            steps += 1;
            if (map.getValue(p)) |other_steps| {
                bestDist = std.math.min(bestDist, p.manhattan_dist(Pos{}));
                bestSteps = std.math.min(bestSteps, steps + other_steps);
            }
        }
    }

    std.debug.warn("{}\n", bestDist);
    std.debug.warn("{}\n", bestSteps);
}
