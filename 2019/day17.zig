const std = @import("std");
const utils = @import("utils.zig");

pub fn main() void {
    var prog = utils.loadProgram("day17_input.txt");

    // part 1
    var out = utils.Machine.init(prog).run(-1) orelse unreachable;
    var stride: usize = 1;
    while (out[stride - 1] != '\n') {
        stride += 1;
    }
    var sum: i64 = 0;
    for (out) |_, i| {
        if (i < stride or i + stride >= out.len) {
            continue;
        }
        var intersection = out[i] == '#' and out[i - 1] == '#' and out[i + 1] == '#' and out[i - stride] == '#' and out[i + stride] == '#';
        if (intersection) {
            var p = utils.Pos.fromStride(i, stride);
            sum += p.x * p.y;
        }
    }
    utils.println(sum);

    // part 2
    prog[0] = 2;
    out = utils.Machine.init(prog).run(([]const u8)(
        \\A,B,A,C,B,C,A,B,A,C
        \\R,6,L,10,R,8,R,8
        \\R,12,L,8,L,10
        \\R,12,L,10,R,6,L,10
        \\n
        \\
    )) orelse unreachable;
    utils.println(out[out.len - 1]);
}
