//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestCountErrors — countErrors 가 ERROR 레벨 진단만 세는지 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCountErrors(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError},
		{Level: diagnostic.LevelWarning},
		{Level: diagnostic.LevelError},
		{Level: diagnostic.LevelWarning},
	}
	if got := countErrors(diags); got != 2 {
		t.Errorf("countErrors = %d, want 2", got)
	}
	if got := countErrors(nil); got != 0 {
		t.Errorf("countErrors(nil) = %d, want 0", got)
	}
}
