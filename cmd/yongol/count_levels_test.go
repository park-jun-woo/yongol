//ff:func feature=cli type=test control=sequence
//ff:what countLevels test — ERROR/WARNING 카운트 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCountLevels(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		errs, warns := countLevels(nil)
		if errs != 0 || warns != 0 {
			t.Fatalf("expected (0,0), got (%d,%d)", errs, warns)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{Level: diagnostic.LevelError},
			{Level: diagnostic.LevelWarning},
			{Level: diagnostic.LevelError},
		}
		errs, warns := countLevels(diags)
		if errs != 2 {
			t.Errorf("expected 2 errors, got %d", errs)
		}
		if warns != 1 {
			t.Errorf("expected 1 warning, got %d", warns)
		}
	})
}
