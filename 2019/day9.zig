const std = @import("std");
const utils = @import("utils.zig");

pub fn main() void {
    var prog = utils.loadProgram("day9_input.txt");

    // part 1
    var out = utils.Machine.init(prog).run(1) orelse unreachable;
    utils.println(out[0]);

    // part 2
    out = utils.Machine.init(prog).run(2) orelse unreachable;
    utils.println(out[0]);
}
