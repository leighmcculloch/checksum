module main

import cli
import io
import os
import crypto.sha256

fn main() {
	mut cmd := cli.Command{
		name: 'checksum'
		description: 'Verify stdin matches SHA256 checksum and pipe to output. Stdin is piped to stdout regardless of verification, but command errors with exit code 1 if verification fails.'
		execute: run
	}
	cmd.add_flag(cli.Flag{
		flag: .string
		required: false
		name: 'checksum'
		abbrev: 'c'
		description: 'SHA256 checksum to verify against stdin'
	})
	cmd.setup()
	cmd.parse(os.args)
}

fn run(cmd cli.Command) ? {
	hash := sha256.new()
	w := io.new_multi_writer(hash, os.stdout())
	io.cp(w, os.stdin()) ?
	checksum := hash.sum([]).hex()
	expected_checksum := cmd.flags.get_string('checksum') or { '' }
	if checksum != expected_checksum {
		return error('sha256 of input $checksum does not match checksum $expected_checksum')
	}
}
