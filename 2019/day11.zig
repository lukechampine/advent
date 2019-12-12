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

const Color = enum {
    Black,
    White,
};

const Panels = std.AutoHashMap(utils.Pos, Color);

const robot = struct {
    const Self = @This();
    a: utils.Actor = undefined,
    m: machine = undefined,
    panels: *Panels = undefined,

    fn run(s: *Self) bool {
        var in: i64 = switch (s.panels.getValue(s.a.pos) orelse Color.Black) {
            Color.Black => i64(0),
            Color.White => 1,
        };
        var color = s.m.run(in) orelse return false;
        _ = s.panels.put(s.a.pos, switch (color) {
            0 => Color.Black,
            1 => Color.White,
            else => unreachable,
        }) catch unreachable;

        var dir = s.m.run(in) orelse return false;
        switch (dir) {
            0 => s.a.turnLeft(),
            1 => s.a.turnRight(),
            else => unreachable,
        }
        s.a.moveForward(1);
        return true;
    }

    fn init(m: machine, panels: *Panels) Self {
        var a = utils.Actor{
            .pos = utils.Pos{ .x = 0, .y = 0 },
            .dir = utils.Dir.Up,
        };
        return Self{ .a = a, .m = m, .panels = panels };
    }
};

pub fn main() void {
    const input = utils.readFile("day11_input.txt");
    var prog = utils.alloc(i64, utils.count(u8, input, ',') + 30000);
    var it = std.mem.separate(input, ",");
    for (prog) |_, i| {
        prog[i] = utils.parseInt(i64, it.next() orelse "0");
    }

    // part 1
    var m = machine.init(prog);
    var panels = utils.makeMap(utils.Pos, Color);
    var r = robot.init(m, &panels);
    while (r.run()) {}
    utils.println(panels.count());

    // part 2
    m = machine.init(prog);
    panels = utils.makeMap(utils.Pos, Color);
    _ = panels.put(utils.Pos{ .x = 0, .y = 0 }, Color.White) catch unreachable;
    r = robot.init(m, &panels);
    while (r.run()) {}
    var minPos = utils.Pos{ .x = 99999, .y = 99999 };
    var maxPos = utils.Pos{ .x = -99999, .y = -99999 };
    var pit = panels.iterator();
    while (pit.next()) |o| {
        minPos.x = std.math.min(minPos.x, o.key.x);
        minPos.y = std.math.min(minPos.y, o.key.y);
        maxPos.x = std.math.max(maxPos.x, o.key.x);
        maxPos.y = std.math.max(maxPos.y, o.key.y);
    }
    var p = utils.Pos{ .y = maxPos.y };
    while (p.y >= minPos.y) : (p.y -= 1) {
        p.x = minPos.x;
        while (p.x <= maxPos.x) : (p.x += 1) {
            var c = panels.getValue(p) orelse Color.Black;
            utils.print(switch (c) {
                Color.Black => " ",
                Color.White => "█",
            });
        }
        utils.print("\n");
    }
}
