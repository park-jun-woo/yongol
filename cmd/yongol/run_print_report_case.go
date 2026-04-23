//ff:func feature=cli type=test-helper control=sequence
//ff:what runPrintReportCase — 단일 printReportCase 를 실행하고 (errors,warnings,err) tuple 을 assert
package main

import (
	"bytes"
	"strings"
	"testing"
)

// runPrintReportCase executes printReport for one table entry and asserts the
// (errors, warnings, err) tuple plus the "validation failed: …" message
// shape. Extracted from TestPrintReportReturnValues so filefunc Q4 (pure
// range body ≤ 10) passes.
func runPrintReportCase(t *testing.T, tc printReportCase) {
	t.Helper()
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
}
