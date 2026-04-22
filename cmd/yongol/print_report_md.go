//ff:func feature=cli type=reporter control=iteration dimension=1 topic=md
//ff:what printReportMD — prints validate.Report as GFM-lite (H2/H3 + bullet + inline code)
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printReportMD renders the validation report in GitHub Flavored Markdown
// (shallow subset): H2 section headers, bullet lists with inline code for
// file locations and rule IDs. No tables — pipe characters render poorly
// in a terminal.
//
// Returns (errors, warnings, err). err is non-nil when any step failed or
// any ERROR diagnostic was emitted; its message embeds both counts:
//
//	validation failed: 3 errors, 1 warnings
func printReportMD(w io.Writer, r *validate.Report) (errors, warnings int, err error) {
	failed := false
	fmt.Fprintln(w, "## Validation")
	fmt.Fprintln(w)
	for _, s := range r.Steps {
		ne, nw := countLevels(s.Diagnostics)
		errors += ne
		warnings += nw
		if s.Status == validate.StatusFail || ne > 0 {
			failed = true
		}
		fmt.Fprintln(w, formatStepLine(s, ne, nw))
	}

	if errors > 0 {
		fmt.Fprintln(w, "\n### Errors")
		for _, s := range r.Steps {
			printLevelDiags(w, s, diagnostic.LevelError)
		}
	}
	if warnings > 0 {
		fmt.Fprintln(w, "\n### Warnings")
		for _, s := range r.Steps {
			printLevelDiags(w, s, diagnostic.LevelWarning)
		}
	}

	fmt.Fprintf(w, "\n%d errors, %d warnings\n", errors, warnings)
	if failed {
		err = fmt.Errorf("validation failed: %d errors, %d warnings", errors, warnings)
	}
	return errors, warnings, err
}
