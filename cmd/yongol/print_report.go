//ff:func feature=cli type=reporter control=selection dimension=1
//ff:what printReport — dispatcher that routes to the appropriate renderer based on the format (md|json|sarif) value
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// Supported output formats for `yongol validate --format`.
const (
	formatMD    = "md"
	formatJSON  = "json"
	formatSARIF = "sarif"
)

// printReport renders a validate.Report in the requested format and returns
// the aggregate ERROR and WARNING counts plus a non-nil err when any step
// failed or any ERROR diagnostic was emitted. The failure message embeds the
// counts so stderr-only CI logs still surface the totals:
//
//	validation failed: 3 errors, 1 warnings
//
// Callers that need the warning count (e.g. generate_cmd gating on
// zero-warning) read the second return value rather than re-walking Report.
//
// Supported formats:
//   - "md" (default, GFM-lite)
//   - "json" (bespoke flat snake_case — PhaseF02)
//   - "sarif" (SARIF 2.1.0 full catalog — PhaseF02)
//
// When specsDir is empty, SARIF / JSON artifactLocation URIs fall back to the
// raw Diagnostic.File value (no rebase).
func printReport(w io.Writer, r *validate.Report, format, specsDir string) (errors, warnings int, err error) {
	switch format {
	case "", formatMD:
		return printReportMD(w, r)
	case formatJSON:
		return printReportJSON(w, r, specsDir)
	case formatSARIF:
		return printReportSARIF(w, r, specsDir)
	default:
		return 0, 0, fmt.Errorf("unknown format %q (supported: md, json, sarif)", format)
	}
}
