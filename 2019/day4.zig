const std = @import("std");
const utils = @import("utils.zig");

fn adj(n: usize) bool {
    var str = utils.formatInt(n);
    var i: usize = 0;
    while (i < str.len - 1) : (i += 1) {
        if (str[i] == str[i + 1]) {
            return true;
        }
    }
    return false;
}

fn monotonic(n: usize) bool {
    var str = utils.formatInt(n);
    var i: usize = 0;
    while (i < str.len - 1) : (i += 1) {
        if (str[i] > str[i + 1]) {
            return false;
        }
    }
    return true;
}

fn adjExactly2(n: usize) bool {
    var str = utils.formatInt(n);
    var i: usize = 0;
    while (i < str.len - 1) : (i += 1) {
        var count: usize = 1;
        while (i < str.len - 1 and str[i] == str[i + 1]) : (i += 1) {
            count += 1;
        }
        if (count == 2) {
            return true;
        }
    }
    return false;
}

pub fn main() void {
    const input = utils.readFile("day4_input.txt");
    const min = utils.parseInt(usize, input[0..6]);
    const max = utils.parseInt(usize, input[7..]);
    // part 1
    var i = min;
    var count: usize = 0;
    while (i < max) : (i += 1) {
        if (adj(i) and monotonic(i))
            count += 1;
    }
    utils.println(count);

    // part 2
    i = min;
    count = 0;
    while (i < max) : (i += 1) {
        if (adjExactly2(i) and monotonic(i))
            count += 1;
    }
    utils.println(count);
}
