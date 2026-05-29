//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCountErrors — 진단 중 ERROR 레벨 개수
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCountErrors(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError},
		{Level: diagnostic.LevelWarning},
		{Level: diagnostic.LevelError},
	}
	if got := countErrors(diags); got != 2 {
		t.Errorf("countErrors = %d, want 2", got)
	}
	if got := countErrors(nil); got != 0 {
		t.Errorf("countErrors(nil) = %d, want 0", got)
	}
}
