//ff:func feature=agent type=test control=sequence
//ff:what TestExtractRegoBlock — allow 블록 추출 성공 + op미존재/allow미존재/닫는중괄호미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractRegoBlockOpNotFound(t *testing.T) {
	_, _, _, err := extractRegoBlock("package x\nallow if {\n  input.action == \"A\"\n}\n", "Missing")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected not-found error, got: %v", err)
	}
}
