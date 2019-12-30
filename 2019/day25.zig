const std = @import("std");
const utils = @import("utils.zig");

fn printASCII(ascii: []i64) void {
    for (ascii) |a| {
        std.debug.warn("{c}", @intCast(u8, a));
    }
}

pub fn main() void {
    var prog = utils.loadProgram("day25_input.txt");
    var m = utils.Machine.init(prog);
    printASCII(m.run(([]const u8)("")) orelse unreachable);
    var buf: [100]u8 = undefined;
    while (true) {
        var line = std.io.readLineSlice(buf[0..]) catch unreachable;
        var input = utils.append(u8, line, '\n');
        printASCII(m.run(input) orelse unreachable);
    }
}
