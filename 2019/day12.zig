const std = @import("std");
const utils = @import("utils.zig");

const Vec3 = struct {
    x: i64 = 0,
    y: i64 = 0,
    z: i64 = 0,

    fn energy(s: Vec3) i64 {
        var sum: i64 = 0;
        inline for (std.meta.fields(Vec3)) |f| {
            sum += utils.abs(@field(s, f.name));
        }
        return sum;
    }
};

const Moon = struct {
    pos: Vec3 = Vec3{},
    vel: Vec3 = Vec3{},

    fn energy(m: Moon) i64 {
        return m.pos.energy() * m.vel.energy();
    }

    fn move(m: *Moon) void {
        inline for (std.meta.fields(Vec3)) |f| {
            @field(m.pos, f.name) += @field(m.vel, f.name);
        }
    }

    fn gravitate(m: *Moon, n: Moon) void {
        inline for (std.meta.fields(Vec3)) |f| {
            @field(m.vel, f.name) += utils.sign(i64, @field(n.pos, f.name) - @field(m.pos, f.name));
        }
    }
};

pub fn main() void {
    const input = utils.readFileLines("day12_input.txt");
    var initPositions: [4]Vec3 = undefined;
    for (initPositions) |*p, i| {
        var coords = utils.splitByte(input[i][0 .. input[i].len - 1], ',');
        inline for (std.meta.fields(Vec3)) |f, j| {
            @field(p, f.name) = utils.parseInt(i64, coords[j][3..]);
        }
    }

    // part 1
    var moons: [4]Moon = undefined;
    for (moons) |*m, i| {
        m.pos = initPositions[i];
        m.vel = Vec3{};
    }
    var steps: usize = 0;
    while (steps < 1000) : (steps += 1) {
        for (moons) |*m, i| {
            for (moons) |n, j| {
                if (i == j) continue;
                m.gravitate(n);
            }
        }
        for (moons) |*m| {
            m.move();
        }
    }
    var energy: i64 = 0;
    for (moons) |m| {
        energy += m.energy();
    }
    utils.println(energy);

    // part 2
    var initialMoons = [4]Moon{
        Moon{ .pos = Vec3{ .x = 17, .y = -7, .z = -11 } },
        Moon{ .pos = Vec3{ .x = 1, .y = 4, .z = -1 } },
        Moon{ .pos = Vec3{ .x = 6, .y = -2, .z = -6 } },
        Moon{ .pos = Vec3{ .x = 19, .y = 11, .z = 9 } },
    };
    moons = initialMoons;
    steps = 0;
    var cycle = Vec3{};
    while (cycle.x * cycle.y * cycle.z == 0) : (steps += 1) {
        inline for (std.meta.fields(Vec3)) |cf| {
            if (@field(cycle, cf.name) == 0) {
                var match: bool = true;
                for (initialMoons) |m, i| {
                    match = match and @field(m.pos, cf.name) == @field(moons[i].pos, cf.name);
                    match = match and @field(m.vel, cf.name) == @field(moons[i].vel, cf.name);
                }
                if (match) {
                    @field(cycle, cf.name) = @intCast(i64, steps);
                }
            }
        }

        for (moons) |*m, i| {
            for (moons) |n, j| {
                if (i == j) continue;
                m.gravitate(n);
            }
        }
        for (moons) |*m| {
            m.move();
        }
    }
    utils.println(utils.lcm(cycle.x, utils.lcm(cycle.y, cycle.z)));
}
