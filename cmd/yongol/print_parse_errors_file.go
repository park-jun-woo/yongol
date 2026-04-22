//ff:func feature=cli type=util control=iteration dimension=1
//ff:what printParseErrorsFile — prints diagnostics belonging to a single file with indentation
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// printParseErrorsFile writes a single file's diagnostics with line info.
func printParseErrorsFile(w io.Writer, file string, diags []diagnostic.Diagnostic) {
	fmt.Fprintf(w, "[%s]\n", file)
	for _, d := range diags {
		line := ""
		if d.Line > 0 {
			line = fmt.Sprintf(" (line %d)", d.Line)
		}
		fmt.Fprintf(w, "  - [%s]%s %s\n", d.Level, line, d.Message)
	}
}
