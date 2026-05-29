//ff:func feature=cli type=test control=sequence
//ff:what TestRunPrintReportCase — runPrintReportCase 헬퍼 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestRunPrintReportCase(t *testing.T) {
	t.Run("AllPass", func(t *testing.T) {
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
	})

	t.Run("OneError", func(t *testing.T) {
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
	})

	t.Run("OneWarning", func(t *testing.T) {
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
	})

	t.Run("ErrorWithMsgCheck", func(t *testing.T) {
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
	})

	t.Run("MultipleSteps", func(t *testing.T) {
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
	})

	t.Run("ErrorNoMsgSubstring", func(t *testing.T) {
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
	})
}
