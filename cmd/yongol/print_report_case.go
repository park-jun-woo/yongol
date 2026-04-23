//ff:type feature=cli type=test-helper
//ff:what printReportCase — TestPrintReportReturnValues 단일 케이스 구조체
package main

import (
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printReportCase is one entry in the TestPrintReportReturnValues table,
// describing the input Report plus the expected return tuple.
type printReportCase struct {
	name         string
	report       *validate.Report
	wantErrors   int
	wantWarnings int
	wantErr      bool
	wantMsgHas   string
}
