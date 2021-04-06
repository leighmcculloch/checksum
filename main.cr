require "digest"
require "admiral"

class App < Admiral::Command
  define_help description: "Pass through checksum checker."
  define_argument checksum : String, long: "checksum", short: "c"

  def run
    d = IO::Digest.new(STDOUT, ::Digest::SHA256.new, IO::Digest::DigestMode::Write)
    IO.copy(STDIN, d)

    b = d.final
    STDERR.puts b.map { |n| "%02x" % (n & 0xFF) }.join
  end
end

App.run
