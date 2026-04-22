//ff:func feature=cli type=util control=iteration dimension=1
//ff:what printParseErrors — sorts and prints diagnostics collected during the parse phase
package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// printParseErrors renders parser-phase diagnostics grouped by file.
func printParseErrors(w io.Writer, diags []diagnostic.Diagnostic) {
	fmt.Fprintln(w, "== Parse Errors ==")
	byFile := map[string][]diagnostic.Diagnostic{}
	for _, d := range diags {
		byFile[d.File] = append(byFile[d.File], d)
	}
	files := make([]string, 0, len(byFile))
	for f := range byFile {
		files = append(files, f)
	}
	sort.Strings(files)
	for _, f := range files {
		printParseErrorsFile(w, f, byFile[f])
	}
	fmt.Fprintf(w, "\n%d parse errors\n", len(diags))
}
