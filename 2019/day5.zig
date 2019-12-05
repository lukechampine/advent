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

fn runMachine(prog: []i32, input: i32) i32 {
    // mutate copy
    var p = utils.dup(i32, prog);
    var i: usize = 0;
    while (i < p.len) {
        var op = @mod(p[i], 100);
        var args = argParser.init(p, i);

        switch (op) {
            // stores
            1 => p[args.imm(3)] = args.get(1) + args.get(2),
            2 => p[args.imm(3)] = args.get(1) * args.get(2),
            7 => p[args.imm(3)] = if (args.get(1) < args.get(2)) i32(1) else 0,
            8 => p[args.imm(3)] = if (args.get(1) == args.get(2)) i32(1) else 0,
            // conditional jumps
            5 => i = if (args.get(1) != 0) @intCast(usize, args.get(2)) else i + 3,
            6 => i = if (args.get(1) == 0) @intCast(usize, args.get(2)) else i + 3,
            // I/O
            3 => p[args.imm(1)] = input,
            4 => if (args.get(1) != 0) return args.get(1),
            else => unreachable,
        }
        i += switch (op) {
            1, 2, 7, 8 => usize(4),
            5, 6 => 0, // already jumped
            3, 4 => 2,
            else => unreachable,
        };
    }
    unreachable;
}

pub fn main() void {
    const input = utils.readFile("day5_input.txt");
    var prog = utils.alloc(i32, utils.count(input, ","));
    var pi: usize = 0;
    var it = std.mem.separate(input, ",");
    while (it.next()) |d| {
        prog[pi] = utils.parseInt(i32, d);
        pi += 1;
    }

    // part 1
    utils.println(runMachine(prog, 1));
    // part 2
    utils.println(runMachine(prog, 5));
}
