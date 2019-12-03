const std = @import("std");
const utils = @import("utils.zig");

fn calcFuel(m: usize) usize {
    var mm = m / 3;
    return mm - std.math.min(mm, 2);
}

fn calcFuelRec(m: usize) usize {
    var mm = calcFuel(m);
    return if (mm == 0) 0 else mm + calcFuelRec(mm);
}

pub fn main() void {
    const input = utils.readFile("day1_input.txt");

    // part 1
    var it = utils.IntIterator(usize).new(input, "\n");
    var sum: usize = 0;
    while (it.next()) |m| {
        sum += calcFuel(m);
    }
    utils.println(sum);

    // part 2
    it = utils.IntIterator(usize).new(input, "\n");
    sum = 0;
    while (it.next()) |m| {
        sum += calcFuelRec(m);
    }
    utils.println(sum);
}
