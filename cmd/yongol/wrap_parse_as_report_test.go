//ff:func feature=cli type=test control=sequence
//ff:what TestWrapParseAsReport — wrapParseAsReport 변환 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestWrapParseAsReport(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError, Message: "bad parse"},
		{Level: diagnostic.LevelError, Message: "another error"},
	}
	r := wrapParseAsReport(diags)
	if len(r.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(r.Steps))
	}
	step := r.Steps[0]
	if step.Name != "parse" {
		t.Errorf("Name = %q, want %q", step.Name, "parse")
	}
	if step.Status != validate.StatusFail {
		t.Errorf("Status = %v, want StatusFail", step.Status)
	}
	if len(step.Diagnostics) != 2 {
		t.Errorf("Diagnostics count = %d, want 2", len(step.Diagnostics))
	}
}
