//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: TestPrintReportReturnValues — printReport 4 cases validating errors/warnings/err return values

package main

import (
	"bytes"
	"strings"
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
	errDiag := func(msg string) diagnostic.Diagnostic {
		return diagnostic.Diagnostic{
			File:    "x.ssac",
			Line:    1,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: msg,
		}
	}
	warnDiag := func(msg string) diagnostic.Diagnostic {
		return diagnostic.Diagnostic{
			File:    "x.ssac",
			Line:    1,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: msg,
		}
	}

	cases := []struct {
		name         string
		report       *validate.Report
		wantErrors   int
		wantWarnings int
		wantErr      bool
		wantMsgHas   string
	}{
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
						errDiag("E-1"), errDiag("E-2"), errDiag("E-3"),
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
						warnDiag("W-1"), warnDiag("W-2"),
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
						errDiag("E-1"),
						warnDiag("W-1"), warnDiag("W-2"), warnDiag("W-3"), warnDiag("W-4"),
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
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			gotErrs, gotWarns, gotErr := printReport(&buf, tc.report, formatMD, "")
			if gotErrs != tc.wantErrors {
				t.Errorf("errors: got %d, want %d", gotErrs, tc.wantErrors)
			}
			if gotWarns != tc.wantWarnings {
				t.Errorf("warnings: got %d, want %d", gotWarns, tc.wantWarnings)
			}
			if tc.wantErr {
				if gotErr == nil {
					t.Fatalf("expected non-nil err, got nil")
				}
				if tc.wantMsgHas != "" && !strings.Contains(gotErr.Error(), tc.wantMsgHas) {
					t.Errorf("err message missing %q: got %q", tc.wantMsgHas, gotErr.Error())
				}
				if !strings.HasPrefix(gotErr.Error(), "validation failed:") {
					t.Errorf("err message must begin with 'validation failed:'; got %q", gotErr.Error())
				}
			} else if gotErr != nil {
				t.Errorf("expected nil err, got %v", gotErr)
			}
		})
	}
}
