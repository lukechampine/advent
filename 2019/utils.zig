const std = @import("std");

pub fn println(n: var) void {
    std.debug.warn("{}\n", n);
}

pub fn alloc(comptime T: type, n: var) []T {
    return std.heap.direct_allocator.alloc(T, n) catch unreachable;
}

pub fn free(n: var) void {
    std.heap.direct_allocator.free(n);
}

pub fn dup(comptime T: type, orig: []T) []T {
    var d = alloc(T, orig.len);
    std.mem.copy(T, d, orig);
    return d;
}

pub fn abs(x: var) @typeOf(x) {
    return std.math.absInt(x) catch unreachable;
}

pub fn makeMap(comptime K: type, comptime V: type) std.AutoHashMap(K, V) {
    return std.AutoHashMap(K, V).init(std.heap.direct_allocator);
}

pub fn count(str: []const u8, sep: []const u8) usize {
    var it = std.mem.separate(str, sep);
    var c: usize = 0;
    while (it.next()) |_| {
        c += 1;
    }
    return c;
}

pub fn readFile(path: []const u8) []const u8 {
    const data = std.io.readFileAlloc(std.heap.direct_allocator, path) catch unreachable;
    return std.mem.trim(u8, data, "\n");
}

pub fn readFileLines(path: []const u8) [][]const u8 {
    const data = readFile(path);
    var num_lines: usize = count(data, "\n");
    var lines = alloc([]const u8, num_lines);
    var it = std.mem.separate(data, "\n");
    num_lines = 0;
    while (it.next()) |line| {
        lines[num_lines] = line;
        num_lines += 1;
    }
    return lines;
}

pub fn parseInt(comptime T: type, str: []const u8) T {
    return std.fmt.parseInt(T, str, 10) catch unreachable;
}

pub const Pos = struct {
    x: i32 = 0,
    y: i32 = 0,

    fn move(p: Pos, dir: u8) Pos {
        return switch (dir) {
            'U' => Pos{ .x = p.x, .y = p.y + 1 },
            'D' => Pos{ .x = p.x, .y = p.y - 1 },
            'L' => Pos{ .x = p.x - 1, .y = p.y },
            'R' => Pos{ .x = p.x + 1, .y = p.y },
            else => unreachable,
        };
    }

    fn manhattan_dist(p: Pos, o: Pos) i32 {
        return abs(p.x - o.x) + abs(p.y - o.y);
    }
};

pub fn IntIterator(comptime T: type) type {
    return struct {
        const Self = @This();
        it: std.mem.SplitIterator,

        pub fn next(self: *Self) ?T {
            const untrimmed = self.it.next() orelse return null;
            const trimmed = std.mem.trim(u8, untrimmed, " ");
            if (trimmed.len == 0) return null;
            return parseInt(T, trimmed);
        }

        pub fn new(in: []const u8, sep: []const u8) Self {
            return Self{
                .it = std.mem.separate(in, sep),
            };
        }
    };
}
