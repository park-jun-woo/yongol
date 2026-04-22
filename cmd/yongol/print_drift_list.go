//ff:func feature=cli type=reporter control=iteration dimension=1
//ff:what printDriftList — prints the Contract Drift section from validate/contract.Run results

package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// printDriftList writes a "Contract Drift (N)" section grouping
// PRV-01/02 diagnostics by file, so the dashboard reads like a scan
// summary rather than a raw diagnostic stream.
func printDriftList(w io.Writer, diags []diagnostic.Diagnostic) {
	fmt.Fprintf(w, "\nContract Drift (%d)\n", len(diags))
	if len(diags) == 0 {
		return
	}
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
		fmt.Fprintf(w, "  %s\n", f)
		for _, d := range byFile[f] {
			fmt.Fprintf(w, "    %s\n", d.Message)
		}
	}
}
