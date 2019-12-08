const std = @import("std");
const utils = @import("utils.zig");

const argParser = struct {
    const Self = @This();
    p: []i32 = undefined,
    i: usize = undefined,
    flags: [3]u1 = undefined,
    fn get(s: Self, n: usize) i32 {
        var a = s.p[s.i + n];
        return if (s.flags[n - 1] == 1) a else s.p[@intCast(usize, a)];
    }
    fn imm(s: Self, n: usize) usize {
        return @intCast(usize, s.p[s.i + n]);
    }
    fn init(p: []i32, i: usize) Self {
        var s = Self{ .p = p, .i = i };
        for (s.flags) |_, fi| {
            const exp = std.math.powi(usize, 10, fi + 2) catch unreachable;
            s.flags[fi] = @intCast(u1, (@intCast(usize, p[i]) / exp) % 10);
        }
        return s;
    }
};

const machine = struct {
    const Self = @This();
    p: []i32 = undefined,
    i: usize = 0,

    fn run(s: *Self, input: i32) ?i32 {
        var usedInput = false;
        while (s.i < s.p.len) {
            var op = @mod(s.p[s.i], 100);
            var args = argParser.init(s.p, s.i);
            switch (op) {
                // stores
                1 => s.p[args.imm(3)] = args.get(1) + args.get(2),
                2 => s.p[args.imm(3)] = args.get(1) * args.get(2),
                7 => s.p[args.imm(3)] = if (args.get(1) < args.get(2)) i32(1) else 0,
                8 => s.p[args.imm(3)] = if (args.get(1) == args.get(2)) i32(1) else 0,
                // conditional jumps
                5 => s.i = if (args.get(1) != 0) @intCast(usize, args.get(2)) else s.i + 3,
                6 => s.i = if (args.get(1) == 0) @intCast(usize, args.get(2)) else s.i + 3,
                // I/O
                3 => {
                    if (usedInput) {
                        return null;
                    }
                    s.p[args.imm(1)] = input;
                    usedInput = true;
                },
                4 => {
                    var out = args.get(1);
                    if (out != 0) {
                        s.i += 2;
                        return out;
                    }
                },
                // halt
                99 => return null,
                else => unreachable,
            }
            s.i += switch (op) {
                1, 2, 7, 8 => usize(4),
                5, 6 => 0, // already jumped
                3, 4 => 2,
                else => unreachable,
            };
        }
        unreachable;
    }

    fn init(p: []i32) Self {
        var s = Self{ .p = utils.dup(i32, p) };
        return s;
    }
};

pub fn main() void {
    const input = utils.readFile("day7_input.txt");
    var prog = utils.alloc(i32, utils.count(u8, input, ',') + 1);
    var it = std.mem.separate(input, ",");
    for (prog) |_, i| {
        prog[i] = utils.parseInt(i32, it.next() orelse unreachable);
    }

    // part 1
    var seqs = utils.perms(i32, 5);
    var max: i32 = -1;
    for (seqs) |seq| {
        var out: i32 = 0;
        for (seq) |s| {
            var m = machine.init(prog);
            _ = m.run(s);
            out = m.run(out) orelse unreachable;
        }
        max = std.math.max(max, out);
    }
    utils.println(max);

    // part 2
    max = -1;
    var ms = utils.alloc(machine, 5);
    for (seqs) |seq| {
        for (seq) |s, i| {
            ms[i] = machine.init(prog);
            _ = ms[i].run(s + 5);
        }
        var out: i32 = 0;
        loop: while (true) {
            for (ms) |*m| {
                out = m.run(out) orelse break :loop;
            }
        }
        max = std.math.max(max, out);
    }
    utils.println(max);
}
