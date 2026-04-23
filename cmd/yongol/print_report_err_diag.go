//ff:func feature=cli type=test-helper control=sequence
//ff:what printReportErrDiag — 테스트용 LevelError Diagnostic 빌더
package main

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// printReportErrDiag returns a ready-to-use LevelError Diagnostic pinned to
// the "x.ssac" / line 1 / PhaseValidate shape used by TestPrintReportReturnValues.
func printReportErrDiag(msg string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    "x.ssac",
		Line:    1,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: msg,
	}
}
