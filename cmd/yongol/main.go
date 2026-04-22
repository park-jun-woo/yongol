//ff:func feature=cli type=command control=sequence
//ff:what main — yongol CLI 엔트리포인트 (cobra 기반)
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

// newRoot builds the yongol root cobra command with every subcommand
// registered. Extracted from main() so integration tests can invoke the
// full CLI tree (SetArgs / SetOut / SetErr) without spawning a child
// process and without taking the os.Exit branch in main().
func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "yongol",
		Short: "Full-stack SSOT orchestrator",
	}

	// Route cobra flag parse errors (unknown flag, bad value) through
	// usageError so main's errors.As branch maps them to exit 2.
	// SetFlagErrorFunc on the root propagates to every subcommand.
	root.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return &usageError{err: err}
	})

	root.AddCommand(versionCmd())
	root.AddCommand(validateCmd())
	root.AddCommand(generateCmd())
	root.AddCommand(chainCmd())
	root.AddCommand(importCmd())
	root.AddCommand(statusCmd())

	return root
}

func main() {
	root := newRoot()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		var ue *usageError
		if errors.As(err, &ue) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
