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

pub fn append(comptime T: type, orig: []T, e: T) []T {
    var d = alloc(T, orig.len + 1);
    std.mem.copy(T, d, orig);
    d[orig.len] = e;
    return d;
}

pub fn deleteSlice(comptime T: type, orig: []T, i: usize) []T {
    std.mem.copy(T, orig[i..], orig[i + 1 ..]);
    return orig[0 .. orig.len - 1];
}

pub fn abs(x: var) @typeOf(x) {
    if (@typeId(@typeOf(x)) == std.builtin.TypeId.ComptimeInt) {
        return if (x > 0) x else -x;
    }
    return if (!@typeOf(x).is_signed) x else std.math.absInt(x) catch unreachable;
}

pub fn gcd(a: var, b: @typeOf(a)) @typeOf(a) {
    var aa = abs(a);
    var ab = abs(b);
    while (ab != 0) {
        var t = aa;
        aa = ab;
        ab = @rem(t, ab);
    }
    return aa;
}

pub fn makeMap(comptime K: type, comptime V: type) std.AutoHashMap(K, V) {
    return std.AutoHashMap(K, V).init(std.heap.direct_allocator);
}

pub fn count(comptime T: type, in: []const T, e: T) usize {
    var c: usize = 0;
    for (in) |x| {
        if (x == e) {
            c += 1;
        }
    }
    return c;
}

pub fn readFile(path: []const u8) []const u8 {
    const data = std.io.readFileAlloc(std.heap.direct_allocator, path) catch unreachable;
    return std.mem.trim(u8, data, "\n");
}

pub fn splitByte(data: []const u8, b: u8) [][]const u8 {
    var sep = [_]u8{b};
    var groups = alloc([]const u8, count(u8, data, b) + 1);
    var it = std.mem.separate(data, sep[0..]);
    for (groups) |_, i| {
        groups[i] = it.next() orelse unreachable;
    }
    return groups;
}

pub fn readFileLines(path: []const u8) [][]const u8 {
    return splitByte(readFile(path), '\n');
}

pub fn replace(comptime T: type, in: []T, old: T, new: T) []T {
    for (in) |e, i| {
        if (e == old) in[i] = new;
    }
    return in;
}

pub fn parseInt(comptime T: type, str: []const u8) T {
    return std.fmt.parseInt(T, str, 10) catch unreachable;
}

pub fn formatInt(buf: []u8, i: var) []const u8 {
    return buf[0..std.fmt.formatIntBuf(buf, i, 10, false, std.fmt.FormatOptions{})];
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

    fn rel(p: Pos, q: Pos) Pos {
        return Pos{ .x = p.x - q.x, .y = p.y - q.y };
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

pub fn perms(comptime T: type, n: T) [][]T {
    if (n == 1) {
        var ps = alloc([]T, 1);
        ps[0] = alloc(T, 1);
        ps[0][0] = 0;
        return ps;
    }
    var ps = perms(T, n - 1);

    // interleave
    var leaved = alloc([]T, ps.len * @intCast(usize, n));
    var li: usize = 0;
    for (ps) |perm| {
        var i: usize = 0;
        while (i <= perm.len) : (i += 1) {
            var withN = alloc(T, perm.len + 1);
            std.mem.copy(T, withN[0..i], perm[0..i]);
            withN[i] = n - 1;
            std.mem.copy(T, withN[i + 1 ..], perm[i..]);
            leaved[li] = withN;
            li += 1;
        }
    }
    return leaved;
}
