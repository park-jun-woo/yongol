//ff:func feature=cli type=reporter control=iteration dimension=1
//ff:what printLevelDiags — prints diagnostics of a given level under a "[step]" header
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printLevelDiags writes diagnostics of the given level under a "[step]" header.
// Multi-line messages and " → Advice: ..." suggestions are indented for readability.
func printLevelDiags(w io.Writer, s validate.StepResult, level diagnostic.Level) {
	first := true
	for _, d := range s.Diagnostics {
		if d.Level != level {
			continue
		}
		if first {
			fmt.Fprintf(w, "[%s]\n", s.Name)
			first = false
		}
		// Prefer the Advice field; fall back to splitting inline advice from Message for legacy compatibility.
		main := d.Message
		advice := d.Advice
		if advice == "" {
			main, advice = splitAdvice(d.Message)
		}
		fmt.Fprintf(w, "  - %s%s\n", formatLocation(d.File, d.Line), main)
		if advice != "" {
			fmt.Fprintf(w, "      ↳ Advice: %s\n", advice)
		}
	}
}
