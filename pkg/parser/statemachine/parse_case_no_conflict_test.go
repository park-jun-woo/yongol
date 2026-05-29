//ff:func feature=statemachine type=parser control=sequence
//ff:what 대소문자 충돌 없는 상태명 정상 파싱 검증

package statemachine

import (
	"testing"
)

func TestParseCaseNoConflict(t *testing.T) {
	content := "```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> open: PublishGig\n    open --> closed: CloseGig\n```"
	_, diags := Parse("test", content, "test.md")
	if len(diags) != 0 {
		t.Errorf("unexpected diagnostics: %v", diags)
	}
}
