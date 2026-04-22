//ff:func feature=cli type=reporter control=sequence topic=json
//ff:what printReportJSON — validate.Report 를 yongol bespoke flat JSON 으로 stdout 출력
package main

import (
	"fmt"
	"io"

	jsonreport "github.com/park-jun-woo/yongol/pkg/report/json"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printReportJSON serialises the validate.Report as the yongol flat JSON
// document and writes it to w. Exit-code semantics mirror the md dispatcher:
// non-nil err when any step failed or any ERROR diagnostic fired. Counts are
// derived so `generate`'s warning gate keeps working regardless of format.
//
// summary.checks is populated from the embedded rule catalog size so that
// consumers can tell "catalog intact" from "catalog missing" at a glance.
func printReportJSON(w io.Writer, r *validate.Report, specsDir string) (errors, warnings int, err error) {
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
	checks := 0
	if cat != nil {
		checks = cat.Len()
	}

	data, emitErr := jsonreport.Emit(r, Version, specsDir, checks)
	if emitErr != nil {
		return errors, warnings, fmt.Errorf("emit json: %w", emitErr)
	}
	if _, werr := w.Write(data); werr != nil {
		return errors, warnings, fmt.Errorf("write json: %w", werr)
	}
	fmt.Fprintln(w)

	if failed {
		err = fmt.Errorf("validation failed: %d errors, %d warnings", errors, warnings)
	}
	return errors, warnings, err
}
