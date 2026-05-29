//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: TestPrintReportReturnValues — printReport 4 cases validating errors/warnings/err return values

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestPrintReportReturnValues covers the four outcome combinations of
// printReport: all-clean, errors-only, warnings-only, mixed. Each checks
// the (errors, warnings, err) tuple and — when err is non-nil — asserts
// the failure message embeds the counts in the
// "validation failed: N errors, M warnings" shape that stderr-only CI
// pipelines rely on.
func TestPrintReportReturnValues(t *testing.T) {
	cases := []printReportCase{
		{
			name: "no-fail",
			report: &validate.Report{Steps: []validate.StepResult{
				{Name: "step-a", Status: validate.StatusPass},
			}},
			wantErrors:   0,
			wantWarnings: 0,
			wantErr:      false,
		},
		{
			name: "errors-only",
			report: &validate.Report{Steps: []validate.StepResult{
				{
					Name:   "step-err",
					Status: validate.StatusFail,
					Diagnostics: []diagnostic.Diagnostic{
						printReportErrDiag("E-1"), printReportErrDiag("E-2"), printReportErrDiag("E-3"),
					},
				},
			}},
			wantErrors:   3,
			wantWarnings: 0,
			wantErr:      true,
			wantMsgHas:   "3 errors, 0 warnings",
		},
		{
			name: "warnings-only",
			report: &validate.Report{Steps: []validate.StepResult{
				{
					Name:   "step-warn",
					Status: validate.StatusPass,
					Diagnostics: []diagnostic.Diagnostic{
						printReportWarnDiag("W-1"), printReportWarnDiag("W-2"),
					},
				},
			}},
			wantErrors:   0,
			wantWarnings: 2,
			wantErr:      false,
		},
		{
			name: "mixed",
			report: &validate.Report{Steps: []validate.StepResult{
				{
					Name:   "step-mix",
					Status: validate.StatusFail,
					Diagnostics: []diagnostic.Diagnostic{
						printReportErrDiag("E-1"),
						printReportWarnDiag("W-1"), printReportWarnDiag("W-2"), printReportWarnDiag("W-3"), printReportWarnDiag("W-4"),
					},
				},
			}},
			wantErrors:   1,
			wantWarnings: 4,
			wantErr:      true,
			wantMsgHas:   "1 errors, 4 warnings",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) { runPrintReportCase(t, tc) })
	}
}
