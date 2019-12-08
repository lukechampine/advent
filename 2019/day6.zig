const std = @import("std");
const utils = @import("utils.zig");

fn recOrbits(o: []const u8, orbits: var) usize {
    return if (orbits.getValue(o)) |parent|
        1 + recOrbits(parent, orbits)
    else
        0;
}

fn orbitPath(o: []const u8, orbits: var) [][]const u8 {
    return if (orbits.getValue(o)) |parent|
        utils.append([]const u8, orbitPath(parent, orbits), o)
    else
        utils.alloc([]const u8, 0);
}

pub fn main() void {
    const input = utils.readFileLines("day6_input.txt");

    var orbits = std.StringHashMap([]const u8).init(std.heap.direct_allocator);
    for (input) |line| {
        var xs = utils.splitByte(line, ')');
        _ = orbits.put(xs[1], xs[0]) catch unreachable;
    }

    var n: usize = 0;
    var it = orbits.iterator();
    while (it.next()) |o| {
        n += recOrbits(o.key, orbits);
    }
    utils.println(n);

    var youPath = orbitPath("YOU", orbits);
    var sanPath = orbitPath("SAN", orbits);

    for (youPath) |_, i| {
        if (!std.mem.eql(u8, youPath[i], sanPath[i])) {
            utils.println(youPath[i + 1 ..].len + sanPath[i + 1 ..].len);
            return;
        }
    }
}
