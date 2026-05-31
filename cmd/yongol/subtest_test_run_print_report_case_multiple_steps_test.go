//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseMultipleSteps — MultipleSteps 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseMultipleSteps(t *testing.T) {

	tc := printReportCase{
		name: "MultipleSteps",
		report: &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "[S-01] bad"},
				},
			},
			{
				Name:   "openapi",
				Status: validate.StatusPass,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelWarning, Message: "[O-01] warn"},
					{Level: diagnostic.LevelWarning, Message: "[O-02] warn"},
				},
			},
		}},
		wantErrors:   1,
		wantWarnings: 2,
		wantErr:      true,
		wantMsgHas:   "1 errors, 2 warnings",
	}
	runPrintReportCase(t, tc)

}
