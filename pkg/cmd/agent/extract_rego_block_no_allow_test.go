//ff:func feature=agent type=test control=sequence
//ff:what TestExtractRegoBlock — allow 블록 추출 성공 + op미존재/allow미존재/닫는중괄호미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractRegoBlockNoAllow(t *testing.T) {
	// The action line exists but no preceding "allow if {" header.
	content := "package x\nsomething input.action == \"A\"\n"
	_, _, _, err := extractRegoBlock(content, "A")
	if err == nil || !strings.Contains(err.Error(), "allow if") {
		t.Fatalf("expected allow-if error, got: %v", err)
	}
}
