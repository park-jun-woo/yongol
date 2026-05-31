//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseErrorWithMsgCheck — ErrorWithMsgCheck 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseErrorWithMsgCheck(t *testing.T) {

	tc := printReportCase{
		name: "ErrorWithMsg",
		report: &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "[S-01] bad"},
					{Level: diagnostic.LevelWarning, Message: "[W-01] warn"},
				},
			},
		}},
		wantErrors:   1,
		wantWarnings: 1,
		wantErr:      true,
		wantMsgHas:   "1 errors, 1 warnings",
	}
	runPrintReportCase(t, tc)

}
