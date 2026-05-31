//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseOneError — OneError 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseOneError(t *testing.T) {

	tc := printReportCase{
		name: "OneError",
		report: &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "[S-01] bad"},
				},
			},
		}},
		wantErrors:   1,
		wantWarnings: 0,
		wantErr:      true,
		wantMsgHas:   "1 errors",
	}
	runPrintReportCase(t, tc)

}
