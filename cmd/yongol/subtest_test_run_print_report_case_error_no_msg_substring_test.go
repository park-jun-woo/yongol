//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestRunPrintReportCaseErrorNoMsgSubstring — ErrorNoMsgSubstring 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestRunPrintReportCaseErrorNoMsgSubstring(t *testing.T) {

	// wantMsgHas empty exercises the false branch of the substring check:
	// only the "validation failed:" prefix is asserted, not a substring.
	tc := printReportCase{
		name: "ErrorNoMsg",
		report: &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "[S-09] boom"},
				},
			},
		}},
		wantErrors:   1,
		wantWarnings: 0,
		wantErr:      true,
		wantMsgHas:   "",
	}
	runPrintReportCase(t, tc)

}
