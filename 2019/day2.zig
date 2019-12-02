const std = @import("std");

fn runMachine(prog: []u32, noun: u32, verb: u32) u32 {
    // mutate copy
    var p = std.heap.direct_allocator.alloc(u32, prog.len) catch unreachable;
    defer std.heap.direct_allocator.free(p);
    std.mem.copy(u32, p, prog);

    p[1] = noun;
    p[2] = verb;
    var i: usize = 0;
    while (i < p.len) : (i += 4) {
        var x = p[i + 1];
        var y = p[i + 2];
        var z = p[i + 3];
        switch (p[i]) {
            1 => {
                p[z] = p[x] + p[y];
            },
            2 => {
                p[z] = p[x] * p[y];
            },
            99 => {
                return p[0];
            },
            else => {},
        }
    }
    unreachable;
}

pub fn main() !void {
    const input = try std.io.readFileAlloc(std.heap.direct_allocator, "day2_input.txt");
    var progBuf: [1000]u32 = undefined;
    var pi: usize = 0;
    var it = std.mem.separate(std.mem.trim(u8, input, "\n"), ",");
    while (it.next()) |d| {
        progBuf[pi] = std.fmt.parseInt(u32, d, 10) catch unreachable;
        pi += 1;
    }
    var prog = progBuf[0..pi];

    // part 1
    var output = runMachine(prog, 12, 2);
    std.debug.warn("{}\n", output);

    // part 2
    var noun: u32 = 0;
    while (noun < 100) : (noun += 1) {
        var verb: u32 = 0;
        while (verb < 100) : (verb += 1) {
            output = runMachine(prog, noun, verb);
            if (output == 19690720) {
                std.debug.warn("{}\n", 100 * noun + verb);
                return;
            }
        }
    }
}
