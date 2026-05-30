//ff:func feature=agent type=test control=iteration dimension=2
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

func TestExtractRegoBlockOpNotFound(t *testing.T) {
	_, _, _, err := extractRegoBlock("package x\nallow if {\n  input.action == \"A\"\n}\n", "Missing")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected not-found error, got: %v", err)
	}
}

func TestExtractRegoBlockNoAllow(t *testing.T) {
	// The action line exists but no preceding "allow if {" header.
	content := "package x\nsomething input.action == \"A\"\n"
	_, _, _, err := extractRegoBlock(content, "A")
	if err == nil || !strings.Contains(err.Error(), "allow if") {
		t.Fatalf("expected allow-if error, got: %v", err)
	}
}

func TestExtractRegoBlockNoClosingBrace(t *testing.T) {
	// "allow if {" opens a brace that is never closed.
	content := "allow if {\n  input.action == \"A\"\n"
	_, _, _, err := extractRegoBlock(content, "A")
	if err == nil || !strings.Contains(err.Error(), "closing brace") {
		t.Fatalf("expected closing-brace error, got: %v", err)
	}
}
