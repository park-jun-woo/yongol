//ff:func feature=agent type=test control=sequence
//ff:what TestExtractRegoBlock — allow 블록 추출 성공 + op미존재/allow미존재/닫는중괄호미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractRegoBlockNoClosingBrace(t *testing.T) {
	// "allow if {" opens a brace that is never closed.
	content := "allow if {\n  input.action == \"A\"\n"
	_, _, _, err := extractRegoBlock(content, "A")
	if err == nil || !strings.Contains(err.Error(), "closing brace") {
		t.Fatalf("expected closing-brace error, got: %v", err)
	}
}
