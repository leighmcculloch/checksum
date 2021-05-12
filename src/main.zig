const std = @import("std");
const io = std.io;
const sha2 = std.crypto.hash.sha2;
const fmt = std.fmt;

pub fn main() anyerror!void {
    const stdout = io.getStdOut();
    const stderr = io.getStdErr();
    const stdin = io.getStdIn();

    var h = sha2.Sha256.init(.{});
    var hw = io.Writer(*@TypeOf(h), error{}, struct {
        fn write(self: *sha2.Sha256, m: []const u8) !usize {
            self.update(m);
            return m.len;
        }
    }.write){ .context = &h };

    var m = io.multiWriter(.{ hw, stdout.writer() });
    try copy(m.writer(), stdin.reader());

    var sum: [sha2.Sha256.digest_length]u8 = undefined;
    h.final(&sum);
    try stderr.writer().print("{s}", .{fmt.fmtSliceHexLower(sum[0..])});
}

fn copy(w: anytype, r: anytype) !void {
    var buf: [1024]u8 = undefined;
    while (true) {
        const nr = try r.read(buf[0..]);
        if (nr == 0) {
            break;
        }
        var i: usize = 0;
        while (i < nr) {
            const nw = try w.write(buf[i..nr]);
            i += nw;
        }
    }
}
