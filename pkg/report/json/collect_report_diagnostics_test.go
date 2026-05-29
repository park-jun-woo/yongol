//ff:func feature=report type=test control=iteration dimension=2 topic=json
//ff:what TestCollectReportDiagnostics — nil 리포트 no-op + 다중 step 누적 검증
package json

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestCollectReportDiagnostics_Nil covers the nil-report guard: nothing is
// accumulated.
func TestCollectReportDiagnostics_Nil(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}
	collectReportDiagnostics(doc, nil, "", "")
	if len(doc.Diagnostics) != 0 {
		t.Errorf("nil report should add nothing, got %d", len(doc.Diagnostics))
	}
	if doc.Summary.Errors != 0 || doc.Summary.Warnings != 0 {
		t.Errorf("counts should be zero, got %+v", doc.Summary)
	}
}

// TestCollectReportDiagnostics_MultiStep covers the nested iteration over
// steps × diagnostics, with only ERROR/WARNING surviving the per-diagnostic
// filter inside appendDiagnostic.
func TestCollectReportDiagnostics_MultiStep(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}
	report := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{File: "a.ssac", Level: diagnostic.LevelError, Message: "[S-1] boom"},
				{File: "b.ssac", Level: diagnostic.Level("INFO"), Message: "[S-2] noise"},
			},
		},
		{
			Name:   "openapi",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{File: "c.yaml", Level: diagnostic.LevelWarning, Message: "[X-3] warn"},
			},
		},
	}}

	collectReportDiagnostics(doc, report, "", "")

	if doc.Summary.Errors != 1 || doc.Summary.Warnings != 1 {
		t.Errorf("counts: %+v, want 1 error / 1 warning", doc.Summary)
	}
	if len(doc.Diagnostics) != 2 {
		t.Fatalf("diagnostics: got %d, want 2 (INFO dropped)", len(doc.Diagnostics))
	}
	if doc.Diagnostics[0].RuleID != "S-1" || doc.Diagnostics[1].RuleID != "X-3" {
		t.Errorf("unexpected order/ids: %+v", doc.Diagnostics)
	}
}
