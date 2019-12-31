const std = @import("std");

pub fn print(n: var) void {
    std.debug.warn("{}", .{n});
}

pub fn println(n: var) void {
    std.debug.warn("{}\n", .{n});
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

pub fn abs(x: var) @TypeOf(x) {
    if (@typeId(@TypeOf(x)) == std.builtin.TypeId.ComptimeInt) {
        return if (x > 0) x else -x;
    }
    return if (!@TypeOf(x).is_signed) x else std.math.absInt(x) catch unreachable;
}

pub fn gcd(a: var, b: @TypeOf(a)) @TypeOf(a) {
    var aa = abs(a);
    var ab = abs(b);
    while (ab != 0) {
        var t = aa;
        aa = ab;
        ab = @rem(t, ab);
    }
    return aa;
}

pub fn lcm(a: var, b: @TypeOf(a)) @TypeOf(a) {
    return @divExact(a * b, gcd(a, b));
}

pub fn makeMap(comptime K: type, comptime V: type) std.AutoHashMap(K, V) {
    return std.AutoHashMap(K, V).init(std.heap.direct_allocator);
}

pub fn makeStringMap(comptime V: type) std.StringHashMap(V) {
    return std.StringHashMap(V).init(std.heap.direct_allocator);
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

pub const Dir = enum {
    Up,
    Down,
    Left,
    Right,

    fn turnLeft(d: Dir) Dir {
        return switch (d) {
            Dir.Up => Dir.Left,
            Dir.Down => Dir.Right,
            Dir.Left => Dir.Down,
            Dir.Right => Dir.Up,
        };
    }

    fn turnRight(d: Dir) Dir {
        return switch (d) {
            Dir.Up => Dir.Right,
            Dir.Down => Dir.Left,
            Dir.Left => Dir.Up,
            Dir.Right => Dir.Down,
        };
    }

    pub fn fromUDLR(d: u8) Dir {
        return switch (d) {
            'U' => Dir.Up,
            'D' => Dir.Down,
            'L' => Dir.Left,
            'R' => Dir.Right,
            else => unreachable,
        };
    }
};

pub const Pos = struct {
    x: i64 = 0,
    y: i64 = 0,

    fn move(p: Pos, dir: Dir) Pos {
        return p.moveN(dir, 1);
    }
    fn moveN(p: Pos, dir: Dir, steps: i64) Pos {
        return switch (dir) {
            Dir.Up => Pos{ .x = p.x, .y = p.y + steps },
            Dir.Down => Pos{ .x = p.x, .y = p.y - steps },
            Dir.Left => Pos{ .x = p.x - steps, .y = p.y },
            Dir.Right => Pos{ .x = p.x + steps, .y = p.y },
        };
    }

    fn rel(p: Pos, q: Pos) Pos {
        return Pos{ .x = p.x - q.x, .y = p.y - q.y };
    }

    fn manhattan_dist(p: Pos, o: Pos) i64 {
        return abs(p.x - o.x) + abs(p.y - o.y);
    }

    pub fn fromStride(i: usize, stride: usize) Pos {
        return Pos{ .x = @intCast(i64, i % stride), .y = @intCast(i64, i / stride) };
    }
};

pub const Actor = struct {
    const Self = @This();
    pos: Pos = undefined,
    dir: Dir = undefined,

    fn moveForward(s: *Self, steps: i64) void {
        s.pos = s.pos.moveN(s.dir, steps);
    }

    fn turnLeft(s: *Self) void {
        s.dir = s.dir.turnLeft();
    }

    fn turnRight(s: *Self) void {
        s.dir = s.dir.turnRight();
    }

    pub fn init(p: Pos, d: Dir) Self {
        return Self{ .pos = p, .dir = d };
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

pub fn sign(comptime T: type, v: T) T {
    return switch (v) {
        1...std.math.maxInt(T) => 1,
        0 => 0,
        std.math.minInt(T)...-1 => -1,
    };
}

const argParser = struct {
    const Self = @This();
    p: []i64 = undefined,
    i: usize = undefined,
    rel: i64 = undefined,
    flags: [3]u2 = undefined,
    fn get(s: Self, n: usize) *i64 {
        return switch (s.flags[n - 1]) {
            0 => &s.p[@intCast(usize, s.p[s.i + n])],
            1 => &s.p[s.i + n],
            2 => &s.p[@intCast(usize, s.rel + s.p[s.i + n])],
            else => unreachable,
        };
    }
    fn init(p: []i64, i: usize, rel: i64) Self {
        var s = Self{ .p = p, .i = i, .rel = rel };
        for (s.flags) |_, fi| {
            const exp = std.math.powi(usize, 10, fi + 2) catch unreachable;
            s.flags[fi] = @intCast(u2, (@intCast(usize, p[i]) / exp) % 10);
        }
        return s;
    }
};

pub const Machine = struct {
    const Self = @This();
    p: []i64 = undefined,
    i: usize = 0,
    rel: i64 = 0,

    fn run(s: *Self, input: var) ?[]i64 {
        var in: []i64 = switch (@TypeOf(input)) {
            []i64 => input,
            i64, comptime_int => blk: {
                var b = alloc(i64, 1);
                b[0] = input;
                break :blk b;
            },
            []u8, []const u8 => blk: {
                var b = alloc(i64, input.len);
                for (b) |_, i| {
                    b[i] = @intCast(i64, input[i]);
                }
                break :blk b;
            },
            else => {
                println(@typeName(@TypeOf(input)));
                unreachable;
            },
        };

        var inputIndex: usize = 0;
        var output: ?[]i64 = null;
        while (s.i < s.p.len) {
            var op = @mod(s.p[s.i], 100);
            var args = argParser.init(s.p, s.i, s.rel);
            switch (op) {
                // stores
                1 => args.get(3).* = args.get(1).* + args.get(2).*,
                2 => args.get(3).* = args.get(1).* * args.get(2).*,
                7 => args.get(3).* = if (args.get(1).* < args.get(2).*) 1 else 0,
                8 => args.get(3).* = if (args.get(1).* == args.get(2).*) 1 else 0,
                // conditional jumps
                5 => s.i = if (args.get(1).* != 0) @intCast(usize, args.get(2).*) else s.i + 3,
                6 => s.i = if (args.get(1).* == 0) @intCast(usize, args.get(2).*) else s.i + 3,
                // I/O
                3 => {
                    if (inputIndex == in.len) {
                        return output;
                    }
                    args.get(1).* = in[inputIndex];
                    inputIndex += 1;
                },
                4 => {
                    var out = args.get(1).*;
                    if (output) |o| {
                        output = append(i64, o, out);
                    } else {
                        var o = alloc(i64, 1);
                        o[0] = out;
                        output = o;
                    }
                },
                // rel
                9 => s.rel += args.get(1).*,
                // halt
                99 => return output,
                else => unreachable,
            }
            s.i += switch (op) {
                1, 2, 7, 8 => @intCast(usize, 4),
                5, 6 => 0, // already jumped
                3, 4, 9 => 2,
                else => unreachable,
            };
        }
        unreachable;
    }

    pub fn init(p: []i64) Self {
        return Self{ .p = dup(i64, p) };
    }
};

pub fn loadProgram(file: []const u8) []i64 {
    const input = readFile(file);
    var prog = alloc(i64, count(u8, input, ',') + 30000);
    var it = std.mem.separate(input, ",");
    for (prog) |_, i| {
        prog[i] = parseInt(i64, it.next() orelse "0");
    }
    return prog;
}
