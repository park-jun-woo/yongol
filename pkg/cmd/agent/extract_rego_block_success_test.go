//ff:func feature=agent type=test control=sequence
//ff:what TestExtractRegoBlock — allow 블록 추출 성공 + op미존재/allow미존재/닫는중괄호미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractRegoBlockSuccess(t *testing.T) {
	content := strings.Join([]string{
		"package authz",
		"",
		"allow if {",
		`  input.action == "ListUsers"`,
		`  input.role == "admin"`,
		"}",
		"",
		"allow if {",
		`  input.action == "Other"`,
		"}",
	}, "\n")

	block, start, end, err := extractRegoBlock(content, "ListUsers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "ListUsers") || !strings.Contains(block, "allow if {") {
		t.Errorf("block missing expected content: %q", block)
	}
	if strings.Contains(block, "Other") {
		t.Errorf("block should not bleed into the next allow rule: %q", block)
	}
	if start < 0 || end <= start {
		t.Errorf("bad line range: start=%d end=%d", start, end)
	}
}
