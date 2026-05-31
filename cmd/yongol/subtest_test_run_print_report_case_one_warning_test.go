//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseOneWarning — OneWarning 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseOneWarning(t *testing.T) {

	tc := printReportCase{
		name: "OneWarning",
		report: &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "check",
				Status: validate.StatusPass,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelWarning, Message: "[W-01] warn"},
				},
			},
		}},
		wantErrors:   0,
		wantWarnings: 1,
		wantErr:      false,
	}
	runPrintReportCase(t, tc)

}
