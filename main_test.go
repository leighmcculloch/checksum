package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	args := []string{"-c", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}
	stdin := strings.NewReader("hello world")
	stdout := strings.Builder{}
	stderr := strings.Builder{}

	exitCode := run(args, stdin, &stdout, &stderr)

	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "hello world", stdout.String())
	assert.Equal(t, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", stderr.String())
}
