const std = @import("std");
const utils = @import("utils.zig");

fn countLayer(layer: [][]const u8, b: u8) usize {
    var n: usize = 0;
    for (layer) |row| {
        n += utils.count(u8, row, b);
    }
    return n;
}

pub fn main() void {
    const maxX = 25;
    const maxY = 6;
    const input = utils.readFile("day8_input.txt");

    var img = utils.alloc([][]u8, input.len / (maxX * maxY));
    for (img) |_, i| {
        img[i] = utils.alloc([]u8, maxY);
        for (img[i]) |_, j| {
            const start = i * (maxX * maxY) + j * (maxX);
            img[i][j] = input[start .. start + maxX];
        }
    }

    // part 1
    var min: usize = 9999999;
    var minLayer: [][]u8 = undefined;
    for (img) |layer, i| {
        var zeros = countLayer(layer, '0');
        if (zeros < min) {
            min = zeros;
            minLayer = layer;
        }
    }
    utils.println(countLayer(minLayer, '1') * countLayer(minLayer, '2'));

    // part 2
    var top = img[0];
    for (top) |row, y| {
        for (row) |p, x| {
            if (p == '2') {
                // find first non-transparent pixel
                find: for (img) |l| {
                    if (l[y][x] != '2') {
                        row[x] = l[y][x];
                        break :find;
                    }
                }
            }
        }
    }
    for (top) |row| {
        utils.println(utils.replace(u8, row, '0', ' '));
    }
}
