const std = @import("std");
const utils = @import("utils.zig");

var prog: []i64 = undefined;

fn inBeam(x: i64, y: i64) bool {
    var in = [_]i64{ x, y };
    var out = utils.Machine.init(prog).run(in[0..]) orelse unreachable;
    return out[0] == 1;
}

fn fitsSquare(y: i64) bool {
    var x = rightEdge(y);
    return inBeam(x, y) and inBeam(x - 99, y + 99);
}

fn rightEdge(y: i64) i64 {
    var x = y; // heuristic
    while (!inBeam(x, y) or inBeam(x + 1, y)) {
        x += 1;
    }
    return x;
}

pub fn main() void {
    prog = utils.loadProgram("day19_input.txt");

    var n: usize = 0;
    var i: usize = 0;
    while (i < 50 * 50) : (i += 1) {
        var p = utils.Pos.fromStride(i, 50);
        if (inBeam(p.x, p.y)) {
            n += 1;
        }
    }
    utils.println(n);

    // part 2
    // binary search
    var bit: i64 = 1 << 12;
    var y: i64 = bit - 1; // start with all bits set
    while (bit > 0) : (bit >>= 1) {
        if (fitsSquare(y & ~bit)) {
            y &= ~bit; // clear bit if we still fit
        }
    }
    var x = rightEdge(y) - 99;
    utils.println(x * 10000 + y);
}
