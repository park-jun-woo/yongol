//ff:func feature=agent type=test control=sequence
//ff:what TestMergeRegoBlock — allow/if/{ 패턴 누락 거부 및 정상 머지 검증

package agent

import (
	"strings"
	"testing"
)

func TestMergeRegoBlock(t *testing.T) {
	original := "package authz\nallow if {\n  input.action == \"Old\"\n}"
	fixed := "allow if {\n  input.action == \"New\"\n}"

	got, err := mergeRegoBlock(original, 1, 4, fixed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "New") {
		t.Errorf("merged missing new rule: %q", got)
	}

	if _, err := mergeRegoBlock(original, 1, 4, "default deny = true"); err == nil {
		t.Error("expected error for block missing 'allow if {' pattern")
	}
}
