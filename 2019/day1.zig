const std = @import("std");

fn IntIterator(comptime T: type) type {
    return struct {
        const Self = @This();
        it: std.mem.SplitIterator,

        fn next(self: *Self) ?T {
            const untrimmed = self.it.next() orelse return null;
            const trimmed = std.mem.trim(u8, untrimmed, " ");
            if (trimmed.len == 0) return null;
            const x = std.fmt.parseInt(T, trimmed, 10) catch unreachable;
            return x;
        }

        fn new(in: []const u8) Self {
            return Self{
                .it = std.mem.separate(in, "\n"),
            };
        }
    };
}

fn calcFuel(m: u64) u64 {
    var mm = m / 3;
    if (mm <= 2) {
        return 0;
    }
    mm -= 2;
    return mm + calcFuel(mm);
}

pub fn main() !void {
    const input = try std.io.readFileAlloc(std.heap.direct_allocator, "day1_input.txt");

    // part 1
    var it = IntIterator(u64).new(input);
    var sum: u64 = 0;
    while (it.next()) |m| {
        sum += (m / 3) - 2;
    }
    std.debug.warn("{}\n", sum);

    // part 2
    it = IntIterator(u64).new(input);
    sum = 0;
    while (it.next()) |m| {
        sum += calcFuel(m);
    }
    std.debug.warn("{}\n", sum);
}
