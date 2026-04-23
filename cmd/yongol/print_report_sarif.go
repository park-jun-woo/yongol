//ff:func feature=cli type=reporter control=iteration dimension=1 topic=sarif
//ff:what printReportSARIF — prints validate.Report as SARIF 2.1.0 JSON to stdout
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/report/sarif"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printReportSARIF serialises the validate.Report as a SARIF 2.1.0 document
// and writes the JSON to w. The md dispatcher's exit-code semantics are
// preserved: non-nil err when any step failed or any ERROR diagnostic fired.
// Counts are still derived so the `generate` command's warning gate keeps
// working regardless of the chosen format.
//
// PhaseF02: loads the embedded rule catalog and passes it to sarif.Emit so
// tool.driver.rules[] carries the full rulebook (150+ entries) with
// shortDescription / helpUri / defaultConfiguration populated.
func printReportSARIF(w io.Writer, r *validate.Report, specsDir string) (errors, warnings int, err error) {
	failed := false
	for _, s := range r.Steps {
		ne, nw := countLevels(s.Diagnostics)
		errors += ne
		warnings += nw
		if s.Status == validate.StatusFail || ne > 0 {
			failed = true
		}
	}

	cat, loadErr := rulecatalog.Load()
	if loadErr != nil {
		return errors, warnings, fmt.Errorf("load rule catalog: %w", loadErr)
	}

	data, emitErr := sarif.Emit(r, Version, specsDir, cat)
	if emitErr != nil {
		return errors, warnings, fmt.Errorf("emit sarif: %w", emitErr)
	}
	if _, werr := w.Write(data); werr != nil {
		return errors, warnings, fmt.Errorf("write sarif: %w", werr)
	}
	// Trailing newline for friendlier terminal/redirect behaviour.
	fmt.Fprintln(w)

	if failed {
		err = fmt.Errorf("validation failed: %d errors, %d warnings", errors, warnings)
	}
	return errors, warnings, err
}
