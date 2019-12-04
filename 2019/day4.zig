const std = @import("std");
const utils = @import("utils.zig");

fn adj(str: []const u8) bool {
    var i: usize = 0;
    while (i < str.len - 1) : (i += 1) {
        if (str[i] == str[i + 1]) {
            return true;
        }
    }
    return false;
}

fn monotonic(str: []const u8) bool {
    var i: usize = 0;
    while (i < str.len - 1) : (i += 1) {
        if (str[i] > str[i + 1]) {
            return false;
        }
    }
    return true;
}

fn adjExactly2(str: []const u8) bool {
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
    var buf: [100]u8 = undefined;
    var i = min;
    var count: usize = 0;
    while (i < max) : (i += 1) {
        var digits = utils.formatInt(buf[0..], i);
        if (adj(digits) and monotonic(digits))
            count += 1;
    }
    utils.println(count);

    // part 2
    i = min;
    count = 0;
    while (i < max) : (i += 1) {
        var digits = utils.formatInt(buf[0..], i);
        if (adjExactly2(digits) and monotonic(digits))
            count += 1;
    }
    utils.println(count);
}
