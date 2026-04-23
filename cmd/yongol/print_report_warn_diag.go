//ff:func feature=cli type=test-helper control=sequence
//ff:what printReportWarnDiag — 테스트용 LevelWarning Diagnostic 빌더
package main

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// printReportWarnDiag returns a ready-to-use LevelWarning Diagnostic pinned
// to the "x.ssac" / line 1 / PhaseValidate shape used by TestPrintReportReturnValues.
func printReportWarnDiag(msg string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    "x.ssac",
		Line:    1,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: msg,
	}
}
