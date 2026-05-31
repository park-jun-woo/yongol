//ff:func feature=agent type=test-helper control=sequence
//ff:what assertGeneratePromptLayer — buildGeneratePrompt 의 공통 헤더/SSaC/레이어 포함·제외 검증 헬퍼
package agent

import (
	"strings"
	"testing"
)

// assertGeneratePromptLayer builds the generate prompt for layer l and asserts
// the common header, SSaC section, the wantHas marker, and absence of wantSkip.
func assertGeneratePromptLayer(t *testing.T, l layer, op, wantHas, wantSkip string) {
	t.Helper()
	got := buildGeneratePrompt(l, op, "make one", "/v1/workflows", "SSAC")
	if !strings.Contains(got, "Feature: make one") {
		t.Errorf("expected feature header, got:\n%s", got)
	}
	if !strings.Contains(got, "SSaC file (CreateWorkflow.ssac):\nSSAC") {
		t.Errorf("expected SSaC section, got:\n%s", got)
	}
	if !strings.Contains(got, wantHas) {
		t.Errorf("expected %q, got:\n%s", wantHas, got)
	}
	if strings.Contains(got, wantSkip) {
		t.Errorf("did not expect %q, got:\n%s", wantSkip, got)
	}
}
