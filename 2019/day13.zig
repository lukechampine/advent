const std = @import("std");
const utils = @import("utils.zig");

const Tile = enum {
    Empty,
    Wall,
    Block,
    Paddle,
    Ball,
};

const Screen = struct {
    const Self = @This();
    tiles: std.AutoHashMap(utils.Pos, Tile) = undefined,
    score: i64 = undefined,

    fn parse(_out: []i64) Self {
        var s = Self{ .tiles = utils.makeMap(utils.Pos, Tile) };
        var out = _out;
        while (out.len > 0) : (out = out[3..]) {
            var p = utils.Pos{ .x = out[0], .y = out[1] };
            if (p.x == -1 and p.y == 0) {
                s.score = out[2];
            } else {
                _ = s.tiles.put(p, @intToEnum(Tile, @intCast(u3, out[2]))) catch unreachable;
            }
        }
        return s;
    }
};

pub fn main() void {
    var prog = utils.loadProgram("day13_input.txt");

    // part 1
    var screen = Screen.parse(utils.Machine.init(prog).run(0) orelse unreachable);
    var n: usize = 0;
    var mit = screen.tiles.iterator();
    while (mit.next()) |kv| {
        if (kv.value == .Block) n += 1;
    }
    utils.println(n);

    // part 2
    var tilt: i64 = 0;
    var ballX: i64 = undefined;
    var paddleX: i64 = undefined;
    prog[0] = 2;
    var m = utils.Machine.init(prog);
    while (true) {
        screen = Screen.parse(m.run(tilt) orelse break);
        mit = screen.tiles.iterator();
        while (mit.next()) |kv| {
            switch (kv.value) {
                .Paddle => {
                    paddleX = kv.key.x;
                },
                .Ball => {
                    ballX = kv.key.x;
                },
                else => {},
            }
        }
        tilt = utils.sign(i64, ballX - paddleX);
    }
    utils.println(screen.score);
}
