//ff:func feature=agent type=test control=sequence
//ff:what TestInsertRegoBlock — allow/if/{ 패턴 누락 거부 및 파일 끝 추가 검증

package agent

import (
	"strings"
	"testing"
)

func TestInsertRegoBlock(t *testing.T) {
	original := "package authz\n\nallow if {\n  input.action == \"A\"\n}\n"
	newBlock := "allow if {\n  input.action == \"B\"\n}\n"

	got, err := insertRegoBlock(original, newBlock)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "\"A\"") || !strings.Contains(got, "\"B\"") {
		t.Errorf("result missing rules: %q", got)
	}
	if strings.Index(got, "\"A\"") > strings.Index(got, "\"B\"") {
		t.Errorf("new block should be appended last: %q", got)
	}

	if _, err := insertRegoBlock(original, "default allow = false"); err == nil {
		t.Error("expected error for block missing 'allow if {' pattern")
	}
}
