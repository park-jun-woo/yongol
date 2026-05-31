//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestCollectResults — nil 리포트 / 다중 step / non-error,warning 필터 / fired 집합 검증
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestCollectResults_FilterAndFired(t *testing.T) {
	report := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{File: "a.ssac", Level: diagnostic.LevelError, Message: "[S-1] boom"},
				{File: "b.ssac", Level: diagnostic.Level("INFO"), Message: "[S-2] noise"},
				{File: "c.ssac", Level: diagnostic.LevelWarning, Message: "no rule prefix here"},
			},
		},
		{
			Name:   "openapi",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{File: "d.yaml", Level: diagnostic.LevelError, Message: "[X-3] bad"},
			},
		},
	}}

	results, fired := collectResults(report, "", nil)

	// INFO filtered out → 3 results (S-1, prefix-less warning, X-3).
	if len(results) != 3 {
		t.Fatalf("results: got %d, want 3", len(results))
	}
	if results[0].RuleID != "S-1" {
		t.Errorf("results[0].ruleId: got %q, want S-1", results[0].RuleID)
	}
	if results[1].RuleID != "" {
		t.Errorf("results[1].ruleId: got %q, want empty", results[1].RuleID)
	}
	if results[2].RuleID != "X-3" {
		t.Errorf("results[2].ruleId: got %q, want X-3", results[2].RuleID)
	}

	// fired contains only the two non-empty rule ids.
	if len(fired) != 2 {
		t.Fatalf("fired: got %d, want 2 (%v)", len(fired), fired)
	}
	if _, ok := fired["S-1"]; !ok {
		t.Errorf("fired missing S-1: %v", fired)
	}
	if _, ok := fired["X-3"]; !ok {
		t.Errorf("fired missing X-3: %v", fired)
	}
}
