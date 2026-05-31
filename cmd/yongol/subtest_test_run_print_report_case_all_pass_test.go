//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseAllPass — AllPass 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseAllPass(t *testing.T) {

	tc := printReportCase{
		name: "AllPass",
		report: &validate.Report{Steps: []validate.StepResult{
			{Name: "check", Status: validate.StatusPass},
		}},
		wantErrors:   0,
		wantWarnings: 0,
		wantErr:      false,
	}
	runPrintReportCase(t, tc)

}
