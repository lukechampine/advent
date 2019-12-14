const std = @import("std");
const utils = @import("utils.zig");

const argParser = struct {
    const Self = @This();
    p: []i64 = undefined,
    i: usize = undefined,
    rel: i64 = undefined,
    flags: [3]u2 = undefined,
    fn get(s: Self, n: usize) *i64 {
        return switch (s.flags[n - 1]) {
            0 => &s.p[@intCast(usize, s.p[s.i + n])],
            1 => &s.p[s.i + n],
            2 => &s.p[@intCast(usize, s.rel + s.p[s.i + n])],
            else => unreachable,
        };
    }
    fn init(p: []i64, i: usize, rel: i64) Self {
        var s = Self{ .p = p, .i = i, .rel = rel };
        for (s.flags) |_, fi| {
            const exp = std.math.powi(usize, 10, fi + 2) catch unreachable;
            s.flags[fi] = @intCast(u2, (@intCast(usize, p[i]) / exp) % 10);
        }
        return s;
    }
};

const machine = struct {
    const Self = @This();
    p: []i64 = undefined,
    i: usize = 0,
    rel: i64 = 0,

    fn run(s: *Self, input: i64) ?i64 {
        var usedInput = false;
        while (s.i < s.p.len) {
            var op = @mod(s.p[s.i], 100);
            var args = argParser.init(s.p, s.i, s.rel);
            switch (op) {
                // stores
                1 => args.get(3).* = args.get(1).* + args.get(2).*,
                2 => args.get(3).* = args.get(1).* * args.get(2).*,
                7 => args.get(3).* = if (args.get(1).* < args.get(2).*) i64(1) else 0,
                8 => args.get(3).* = if (args.get(1).* == args.get(2).*) i64(1) else 0,
                // conditional jumps
                5 => s.i = if (args.get(1).* != 0) @intCast(usize, args.get(2).*) else s.i + 3,
                6 => s.i = if (args.get(1).* == 0) @intCast(usize, args.get(2).*) else s.i + 3,
                // I/O
                3 => {
                    if (usedInput) {
                        return null;
                    }
                    args.get(1).* = input;
                    usedInput = true;
                },
                4 => {
                    var out = args.get(1).*;
                    s.i += 2;
                    return out;
                },
                // rel
                9 => s.rel += args.get(1).*,
                // halt
                99 => return null,
                else => unreachable,
            }
            s.i += switch (op) {
                1, 2, 7, 8 => usize(4),
                5, 6 => 0, // already jumped
                3, 4, 9 => 2,
                else => unreachable,
            };
        }
        unreachable;
    }

    fn init(p: []i64) Self {
        return Self{ .p = utils.dup(i64, p) };
    }
};

pub fn main() void {
    const input = utils.readFile("day13_input.txt");
    var prog = utils.alloc(i64, utils.count(u8, input, ',') + 30000);
    var it = std.mem.separate(input, ",");
    for (prog) |_, i| {
        prog[i] = utils.parseInt(i64, it.next() orelse "0");
    }

    // part 1
    var m = machine.init(prog);
    var tiles = utils.makeMap(utils.Pos, i64);
    while (true) {
        var x = m.run(0) orelse break;
        var y = m.run(0) orelse break;
        var id = m.run(0) orelse break;
        _ = tiles.put(utils.Pos{ .x = x, .y = y }, id) catch unreachable;
    }
    var n: usize = 0;
    var mit = tiles.iterator();
    while (mit.next()) |kv| {
        if (kv.value == 2) n += 1;
    }
    utils.println(n);

    // part 2
    var tilt: i64 = 0;
    var ballX: i64 = undefined;
    var paddleX: i64 = undefined;
    prog[0] = 2;
    m = machine.init(prog);
    tiles = utils.makeMap(utils.Pos, i64);
    while (true) {
        var x = m.run(tilt) orelse break;
        var y = m.run(tilt) orelse break;
        var id = m.run(tilt) orelse break;
        _ = tiles.put(utils.Pos{ .x = x, .y = y }, id) catch unreachable;

        switch (id) {
            3 => paddleX = x,
            4 => ballX = x,
            else => {},
        }
        tilt = utils.sign(i64, ballX - paddleX);
    }
    var score = tiles.get(utils.Pos{ .x = -1, .y = 0 }) orelse unreachable;
    utils.println(score.value);
}
