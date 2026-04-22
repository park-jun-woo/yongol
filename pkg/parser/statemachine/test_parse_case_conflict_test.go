//ff:func feature=statemachine type=parser control=iteration dimension=1
//ff:what 대소문자 충돌 상태명(draft/Draft) 에러 진단 반환 검증

package statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseCaseConflict(t *testing.T) {
	content := "```mermaid\nstateDiagram-v2\n    [*] --> draft\n    Draft --> open: PublishGig\n    open --> closed: CloseGig\n```"
	_, diags := Parse("test", content, "test.md")
	if len(diags) == 0 {
		t.Error("expected diagnostics for case-conflicting state names draft/Draft")
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
