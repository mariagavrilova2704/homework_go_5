package cliargs

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ErrUsage marks invalid command-line arguments.
var ErrUsage = errors.New("usage: args <name> [repeat]")

// Options contains validated user arguments.
type Options struct {
	Name   string
	Repeat int
}

// UserArgs returns a new slice without os.Args[0].
func UserArgs(all []string) []string {
	// TODO: implement according to docs/task.md.
	if len(all) <= 1 {
		return []string{}
	}
	newsl := make([]string, len(all[1:]))
	copy(newsl, all[1:])
	return newsl
}

// Parse validates the arguments passed after os.Args[0].
func Parse(args []string) (Options, error) {
	// TODO: implement according to docs/task.md.
	if len(args) < 1 || len(args) > 2 {
		return Options{}, fmt.Errorf("error: %w", ErrUsage)
	}

	name := strings.TrimSpace(args[0])
	if name == "" {
		return Options{}, fmt.Errorf("error: %w", ErrUsage)
	}
	if len(args) == 1 {
		return Options{name, 1}, nil
	}
	num, err := strconv.Atoi(args[1])
	if err != nil || num <= 0 {
		return Options{}, fmt.Errorf("repeat must be positive int: %w", ErrUsage)
	}

	return Options{name, num}, nil
}

// Run writes the command result to the supplied writer.
func Run(args []string, out io.Writer) error {
	// TODO: implement according to docs/task.md.

	arg, err := Parse(args)
	if err != nil {
		return err
	}

	for i := 0; i < arg.Repeat; i++ {
		_, err := fmt.Fprintf(out, "hello, %s\n", arg.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

// Example returns deterministic output used by cmd and the example test.
func Example() string {
	var out strings.Builder
	_ = Run([]string{"Maria", "2"}, &out)
	return out.String()
}
