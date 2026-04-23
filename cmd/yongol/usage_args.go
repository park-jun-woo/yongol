//ff:func feature=cli type=util control=sequence
//ff:what usageArgs — cobra.PositionalArgs 검증 실패를 *usageError 로 래핑
package main

import (
	"github.com/spf13/cobra"
)

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
