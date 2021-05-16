package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	exitCode := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	checksum := ""

	flagSet := flag.NewFlagSet("checksum", flag.ExitOnError)
	flagSet.SetOutput(stderr)
	flagSet.StringVar(&checksum, "c", checksum, "sha256 checksum to verify against stdin")
	flagSet.Parse(args)

	h := sha256.New()
	w := io.MultiWriter(h, stdout)
	_, err := io.Copy(w, stdin)
	if err != nil {
		fmt.Fprintf(stderr, "error: %s\n", err)
		return 1
	}
	sum := h.Sum(nil)
	sumHex := hex.EncodeToString(sum)
	if sumHex != checksum {
		fmt.Fprintf(stderr, "error: sha256 of input %s does not match checksum %s\n", sumHex, checksum)
		return 1
	}

	fmt.Fprintln(stderr, sumHex)
	return 0
}
