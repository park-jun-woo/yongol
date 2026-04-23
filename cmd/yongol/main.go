//ff:func feature=cli type=command control=sequence
//ff:what main — yongol CLI entry point (cobra-based)
package main

import (
	"errors"
	"fmt"
	"os"
)

// Version is set at build time via -ldflags.
var Version = "dev"

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
