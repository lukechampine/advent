const std = @import("std");
const utils = @import("utils.zig");

const Color = enum {
    Black,
    White,
};

const Panels = std.AutoHashMap(utils.Pos, Color);

const robot = struct {
    const Self = @This();
    a: utils.Actor = undefined,
    m: utils.Machine = undefined,
    panels: *Panels = undefined,

    fn run(s: *Self) bool {
        var in: i64 = switch (s.panels.getValue(s.a.pos) orelse Color.Black) {
            Color.Black => i64(0),
            Color.White => 1,
        };
        var out = s.m.run(in) orelse return false;
        _ = s.panels.put(s.a.pos, @intToEnum(Color, @intCast(u1, out[0]))) catch unreachable;
        switch (out[1]) {
            0 => s.a.turnLeft(),
            1 => s.a.turnRight(),
            else => unreachable,
        }
        s.a.moveForward(1);
        return true;
    }

    fn init(m: utils.Machine, panels: *Panels) Self {
        var a = utils.Actor{
            .pos = utils.Pos{ .x = 0, .y = 0 },
            .dir = utils.Dir.Up,
        };
        return Self{ .a = a, .m = m, .panels = panels };
    }
};

pub fn main() void {
    var prog = utils.loadProgram("day11_input.txt");

    // part 1
    var m = utils.Machine.init(prog);
    var panels = utils.makeMap(utils.Pos, Color);
    var r = robot.init(m, &panels);
    while (r.run()) {}
    utils.println(panels.count());

    // part 2
    m = utils.Machine.init(prog);
    panels = utils.makeMap(utils.Pos, Color);
    _ = panels.put(utils.Pos{ .x = 0, .y = 0 }, Color.White) catch unreachable;
    r = robot.init(m, &panels);
    while (r.run()) {}
    var minPos = utils.Pos{ .x = 99999, .y = 99999 };
    var maxPos = utils.Pos{ .x = -99999, .y = -99999 };
    var pit = panels.iterator();
    while (pit.next()) |o| {
        minPos.x = std.math.min(minPos.x, o.key.x);
        minPos.y = std.math.min(minPos.y, o.key.y);
        maxPos.x = std.math.max(maxPos.x, o.key.x);
        maxPos.y = std.math.max(maxPos.y, o.key.y);
    }
    var p = utils.Pos{ .y = maxPos.y };
    while (p.y >= minPos.y) : (p.y -= 1) {
        p.x = minPos.x;
        while (p.x <= maxPos.x) : (p.x += 1) {
            var c = panels.getValue(p) orelse Color.Black;
            utils.print(switch (c) {
                Color.Black => " ",
                Color.White => "█",
            });
        }
        utils.print("\n");
    }
}
