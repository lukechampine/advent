const std = @import("std");
const utils = @import("utils.zig");

pub fn main() void {
    var prog = utils.loadProgram("day21_input.txt");

    // part 1
    // D && (!A || !C)
    var out = utils.Machine.init(prog).run(@as([]const u8,
        \\NOT A J
        \\NOT C T
        \\OR T J
        \\AND D J
        \\WALK
        \\
    )) orelse unreachable;
    utils.println(out[out.len - 1]);

    // part 2
    // D && (!A || (!B && H) || (!C && (E || H)))
    out = utils.Machine.init(prog).run(@as([]const u8,
        \\NOT C T
        \\OR E J
        \\OR H J
        \\AND T J
        \\NOT B T
        \\AND H T
        \\OR T J
        \\NOT A T
        \\OR T J
        \\AND D J
        \\RUN
        \\
    )) orelse unreachable;
    utils.println(out[out.len - 1]);
}
