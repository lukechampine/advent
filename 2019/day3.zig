const std = @import("std");
const utils = @import("utils.zig");

pub fn main() !void {
    const lines = utils.readFileLines("day3_input.txt");
    var map = utils.makeMap(utils.Pos, usize);
    var p = utils.Pos{};
    var steps: usize = 0;
    var it = std.mem.separate(lines[0], ",");
    while (it.next()) |d| {
        var n = utils.parseInt(i64, d[1..]);
        while (n > 0) : (n -= 1) {
            p = p.move(utils.Dir.fromUDLR(d[0]));
            steps += 1;
            _ = try map.put(p, steps);
        }
    }

    p = utils.Pos{};
    steps = 0;
    it = std.mem.separate(lines[1], ",");
    var bestDist: i64 = 99999999;
    var bestSteps: usize = 99999999;
    while (it.next()) |d| {
        var n = utils.parseInt(i64, d[1..]);
        while (n > 0) : (n -= 1) {
            p = p.move(utils.Dir.fromUDLR(d[0]));
            steps += 1;
            if (map.getValue(p)) |other_steps| {
                bestDist = std.math.min(bestDist, p.manhattan_dist(utils.Pos{}));
                bestSteps = std.math.min(bestSteps, steps + other_steps);
            }
        }
    }

    utils.println(bestDist);
    utils.println(bestSteps);
}
