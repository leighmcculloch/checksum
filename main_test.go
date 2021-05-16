package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun_success(t *testing.T) {
	args := []string{"-c", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}
	stdin := strings.NewReader("hello world")
	stdout := strings.Builder{}
	stderr := strings.Builder{}

	exitCode := run(args, stdin, &stdout, &stderr)

	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "hello world", stdout.String())
	assert.Equal(t, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9\n", stderr.String())
}

func TestRun_failWrongChecksum(t *testing.T) {
	args := []string{"-c", "0000000000000000000000000000000000000000000000000000000000000000"}
	stdin := strings.NewReader("hello world")
	stdout := strings.Builder{}
	stderr := strings.Builder{}

	exitCode := run(args, stdin, &stdout, &stderr)

	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "hello world", stdout.String())
	assert.Equal(t, "error: sha256 of input b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9 does not match checksum 0000000000000000000000000000000000000000000000000000000000000000\n", stderr.String())
}

func TestRun_failNoChecksum(t *testing.T) {
	args := []string{}
	stdin := strings.NewReader("hello world")
	stdout := strings.Builder{}
	stderr := strings.Builder{}

	exitCode := run(args, stdin, &stdout, &stderr)

	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "hello world", stdout.String())
	assert.Equal(t, "error: sha256 of input b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9 does not match checksum \n", stderr.String())
}

