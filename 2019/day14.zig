const std = @import("std");
const utils = @import("utils.zig");

const chemical = struct {
    name: []const u8,
    amount: usize,

    fn fromStr(str: []const u8) chemical {
        var parts = utils.splitByte(std.mem.trim(u8, str, " "), ' ');
        return chemical{
            .name = parts[1],
            .amount = utils.parseInt(usize, parts[0]),
        };
    }
};

const reaction = struct {
    inputs: []chemical,
    output: chemical,
};

fn produce(oc: chemical, excess: *std.StringHashMap(usize), rs: std.StringHashMap(reaction)) usize {
    var c = oc;
    var e = std.math.min(c.amount, excess.getValue(c.name) orelse 0);
    if (e > 0) {
        (excess.get(c.name) orelse unreachable).value -= e;
        c.amount -= e;
    }
    if (c.amount == 0 or std.mem.eql(u8, c.name, "ORE")) {
        return c.amount;
    }
    // lookup inputs and reaction multiplier
    var r = rs.getValue(c.name) orelse unreachable;
    var m = c.amount / r.output.amount;
    if (c.amount % r.output.amount != 0) {
        m += 1;
    }
    // produce each input, summing ore required
    var ore: usize = 0;
    for (r.inputs) |in| {
        var mi = chemical{
            .name = in.name,
            .amount = in.amount * m,
        };
        ore += produce(mi, excess, rs);
    }
    var gop = excess.getOrPut(c.name) catch unreachable;
    if (!gop.found_existing) {
        gop.kv.value = 0;
    }
    gop.kv.value += (m * r.output.amount) - c.amount;
    return ore;
}

pub fn main() void {
    const input = utils.readFileLines("day14_input.txt");

    var reactions = utils.makeStringMap(reaction);
    for (input) |line| {
        var chems = utils.splitByte(line, '=');
        var ins = utils.splitByte(chems[0], ',');
        var output = chemical.fromStr(chems[1][2..]);
        var inputs = utils.alloc(chemical, ins.len);
        for (ins) |in, i| {
            inputs[i] = chemical.fromStr(in);
        }
        var r = reaction{ .inputs = inputs, .output = output };
        _ = reactions.put(r.output.name, r) catch unreachable;
    }

    // part 1
    var excess = utils.makeStringMap(usize);
    var goal = chemical{ .name = "FUEL", .amount = 1 };
    utils.println(produce(goal, &excess, reactions));

    // part 2
    var fuel: usize = 0;
    var interval: usize = 10000000; // arbitrary
    while (interval > 0) : (interval /= 2) {
        excess.clear();
        _ = excess.put("ORE", 1000000000000) catch unreachable;
        goal.amount = fuel;
        var ore = produce(goal, &excess, reactions);
        fuel = if (ore == 0) fuel + interval else fuel - interval;
    }
    utils.println(fuel);
}
