package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	exitCode := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	cmd := cobra.Command{
		Use:     "checksum",
		Long:    "Verify stdin matches SHA256 checksum and pipe to output. Stdin is piped to stdout regardless of verification, but command errors with exit code 1 if verification fails.",
		Example: "curl ... | checksum -c <hash> | tar xz -C <dir>",
		SilenceUsage: true,
	}
	cmd.SetIn(stdin)
	cmd.SetOutput(stdout)
	cmd.SetErr(stderr)

	checksum := ""
	cmd.Flags().StringVarP(&checksum, "checksum", "c", checksum, "SHA256 checksum to verify against stdin")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		h := sha256.New()

		w := io.MultiWriter(h, cmd.OutOrStdout())
		_, err := io.Copy(w, cmd.InOrStdin())
		if err != nil {
			return err
		}

		sum := h.Sum(nil)
		sumHex := hex.EncodeToString(sum)
		if sumHex != checksum {
			return fmt.Errorf("sha256 of input %s does not match checksum %s", sumHex, checksum)
		}

		fmt.Fprint(cmd.ErrOrStderr(), sumHex)
		return nil
	}

	err := cmd.Execute()
	if err != nil {
		return 1
	}
	return 0
}
