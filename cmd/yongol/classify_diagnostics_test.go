//ff:func feature=cli type=test control=sequence
//ff:what classifyDiagnostics test — ERROR/WARNING 분류 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestClassifyDiagnostics(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		errs, warns := classifyDiagnostics(nil)
		if len(errs) != 0 || len(warns) != 0 {
			t.Fatalf("expected empty, got errs=%d warns=%d", len(errs), len(warns))
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{Level: diagnostic.LevelError, Message: "e1"},
			{Level: diagnostic.LevelWarning, Message: "w1"},
			{Level: diagnostic.LevelError, Message: "e2"},
			{Level: diagnostic.LevelWarning, Message: "w2"},
			{Level: diagnostic.LevelWarning, Message: "w3"},
		}
		errs, warns := classifyDiagnostics(diags)
		if len(errs) != 2 {
			t.Errorf("expected 2 errors, got %d", len(errs))
		}
		if len(warns) != 3 {
			t.Errorf("expected 3 warnings, got %d", len(warns))
		}
	})
}
