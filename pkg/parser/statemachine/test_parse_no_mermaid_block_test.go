//ff:func feature=statemachine type=parser control=iteration dimension=1
//ff:what mermaid 블록 없는 입력에서 에러 진단 반환 검증

package statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseNoMermaidBlock(t *testing.T) {
	_, diags := Parse("test", "# No mermaid block here", "test.md")
	if len(diags) == 0 {
		t.Error("expected diagnostics for missing mermaid block")
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected ERROR level, got %s", d.Level)
		}
		if d.Phase != diagnostic.PhaseParse {
			t.Errorf("expected parse phase, got %s", d.Phase)
		}
	}
}
