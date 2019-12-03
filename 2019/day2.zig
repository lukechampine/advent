const std = @import("std");
const utils = @import("utils.zig");

fn runMachine(prog: []u32, noun: u32, verb: u32) u32 {
    // mutate copy
    var p = utils.dup(u32, prog);
    p[1] = noun;
    p[2] = verb;
    var i: usize = 0;
    while (i < p.len and p[i] != 99) : (i += 4) {
        var x = p[i + 1];
        var y = p[i + 2];
        var z = p[i + 3];
        p[z] = switch (p[i]) {
            1 => p[x] + p[y],
            2 => p[x] * p[y],
            else => p[z],
        };
    }
    return p[0];
}

pub fn main() void {
    const input = utils.readFile("day2_input.txt");
    var prog = utils.alloc(u32, utils.count(input, ","));
    var pi: usize = 0;
    var it = std.mem.separate(input, ",");
    while (it.next()) |d| {
        prog[pi] = utils.parseInt(u32, d);
        pi += 1;
    }

    // part 1
    var output = runMachine(prog, 12, 2);
    utils.println(output);

    // part 2
    var noun: u32 = 0;
    while (noun < 100) : (noun += 1) {
        var verb: u32 = 0;
        while (verb < 100) : (verb += 1) {
            output = runMachine(prog, noun, verb);
            if (output == 19690720) {
                utils.println(100 * noun + verb);
                return;
            }
        }
    }
}
