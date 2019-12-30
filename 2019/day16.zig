const std = @import("std");
const utils = @import("utils.zig");

fn phase(vec: []i64) []i64 {
    var prod = utils.alloc(i64, vec.len);
    var stride: usize = 1;
    while (stride <= vec.len) : (stride += 1) {
        var sum: i64 = 0;
        var j = stride - 1;
        while (j < vec.len) {
            var k: usize = 0;
            while (k < stride and j < vec.len) : (k += 1) {
                sum += vec[j];
                j += 1;
            }
            j += stride;
            k = 0;
            while (k < stride and j < vec.len) : (k += 1) {
                sum -= vec[j];
                j += 1;
            }
            j += stride;
        }
        prod[stride - 1] = @rem(utils.abs(sum), 10);
    }
    return prod;
}

fn abs(x: i64) i64 {
    return std.math.absInt(x) catch unreachable;
}

fn addPhase(off: usize, vec: []i64) void {
    var i: usize = vec.len - 2;
    while (i >= off - 1) : (i -= 1) {
        vec[i] += vec[i + 1];
        vec[i + 1] = @rem(abs(vec[i + 1]), 10);
    }
}

fn printDigits(ds: []i64) void {
    for (ds) |c| {
        utils.print(c);
    }
    utils.print("\n");
}

pub fn main() void {
    const input = utils.readFile("day16_input.txt");
    var ds = utils.alloc(i64, input.len);
    for (ds) |_, i| {
        ds[i] = utils.parseInt(i64, input[i .. i + 1]);
    }
    // part 1
    var p: usize = 0;
    while (p < 100) : (p += 1) {
        ds = phase(ds);
    }
    printDigits(ds[0..8]);

    // part 2
    var ds10k = utils.alloc(i64, ds.len * 10000);
    var offset = utils.parseInt(usize, input[0..7]);
    for (ds10k) |_, i| {
        var di = i % ds.len;
        ds10k[i] = utils.parseInt(i64, input[di .. di + 1]);
    }
    p = 0;
    while (p < 100) : (p += 1) {
        addPhase(offset, ds10k);
    }
    printDigits(ds10k[offset .. offset + 8]);
}
