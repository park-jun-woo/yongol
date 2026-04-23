//ff:func feature=cli type=model control=sequence
//ff:what usageError — cobra usage failure wrapper; main maps it to exit code 2

package main

import (
	"github.com/spf13/cobra"
)

// usageError wraps any error that should produce exit code 2 (POSIX usage
// error convention). Uses include cobra Args validation failures and
// unknown-flag errors routed via FlagErrorFunc.
type usageError struct{ err error }

// Error returns the wrapped error message.
func (u *usageError) Error() string { return u.err.Error() }

// Unwrap exposes the underlying error for errors.Is / errors.As.
func (u *usageError) Unwrap() error { return u.err }

// usageArgs wraps a cobra.PositionalArgs validator so its failure surfaces
// as *usageError. main.go's errors.As branch then maps it to exit code 2.
func usageArgs(a cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if err := a(cmd, args); err != nil {
			return &usageError{err: err}
		}
		return nil
	}
}
