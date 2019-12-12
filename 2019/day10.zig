const std = @import("std");
const utils = @import("utils.zig");

fn strideTowards(p: utils.Pos, q: utils.Pos) utils.Pos {
    if (std.meta.eql(p, q)) {
        return q;
    }
    var dx = q.x - p.x;
    var dy = q.y - p.y;
    var gcd = utils.gcd(utils.abs(dx), utils.abs(dy));
    return utils.Pos{
        .x = p.x + @divTrunc(dx, gcd),
        .y = p.y + @divTrunc(dy, gcd),
    };
}

pub fn main() void {
    const lines = utils.readFileLines("day10_input.txt");
    var asteroids = (&[_]utils.Pos{})[0..];
    for (lines) |line, y| {
        for (line) |c, x| {
            if (c == '#') {
                asteroids = utils.append(utils.Pos, asteroids, utils.Pos{
                    .x = @intCast(i32, x),
                    .y = @intCast(i32, y),
                });
            }
        }
    }

    var max: usize = 0;
    var maxIndex: usize = undefined;
    for (asteroids) |a, i| {
        var n: usize = 0;
        next: for (asteroids) |b| {
            if (std.meta.eql(a, b)) continue;

            var p = strideTowards(a, b);
            while (!std.meta.eql(p, b)) {
                if (lines[@intCast(usize, p.y)][@intCast(usize, p.x)] == '#') {
                    continue :next;
                }
                p = strideTowards(p, b);
            }
            n += 1;
        }
        if (n > max) {
            max = n;
            maxIndex = i;
        }
    }
    utils.println(max);

    var station = asteroids[maxIndex];
    asteroids = utils.deleteSlice(utils.Pos, asteroids, maxIndex);

    var precedences = utils.makeMap(utils.Pos, usize);
    for (asteroids) |a| {
        var n: usize = 0;
        var p = strideTowards(a, station);
        while (!std.meta.eql(p, station)) {
            if (lines[@intCast(usize, p.y)][@intCast(usize, p.x)] == '#') {
                n += 1;
            }
            p = strideTowards(p, station);
        }
        _ = precedences.put(a, n) catch unreachable;
    }

    const asteroidWithRadians = struct {
        a: utils.Pos,
        r: f64,
    };
    var awrs = utils.alloc(asteroidWithRadians, asteroids.len);
    for (asteroids) |a, i| {
        awrs[i].a = a;

        var r = a.rel(station);
        var phi = std.math.atan2(f64, @intToFloat(f64, r.y), @intToFloat(f64, r.x));
        // rotate polar axis to point straight "up"
        if (phi < -std.math.pi / 2.0) {
            phi += 2 * std.math.pi;
        }
        // add another full rotation for each asteroid "in the way"
        var p = precedences.getValue(a) orelse unreachable;
        phi += 2 * std.math.pi * @intToFloat(f64, p);
        awrs[i].r = phi;
    }

    var sortFn = struct {
        fn f(a: asteroidWithRadians, b: asteroidWithRadians) bool {
            return a.r < b.r;
        }
    }.f;
    std.sort.sort(asteroidWithRadians, awrs, sortFn);
    var a = awrs[199].a;
    utils.println(a.x * 100 + a.y);
}
