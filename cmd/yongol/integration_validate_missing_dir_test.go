//ff:func feature=cli type=test control=sequence
//ff:what validate missing-dir — 존재하지 않는 디렉토리 시 detect SSOTs 에러 (exit 1)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationValidate_MissingDir verifies exit 1 + "detect SSOTs"
// error when specs-dir does not exist on disk. Guards PhaseC01's
// 1-vs-2 exit-code distinction: this is a runtime error, not a usage one.
func TestIntegrationValidate_MissingDir(t *testing.T) {
	_, _, err := runCmd(t, "validate", "/does/not/exist/yongol-test")
	if err == nil {
		t.Fatal("expected error for missing specs-dir, got nil")
	}
	if isUsageError(err) {
		t.Fatalf("missing dir should NOT be *usageError (exit 1, not 2): %v", err)
	}
	if !strings.Contains(err.Error(), "detect SSOTs") {
		t.Errorf("expected error to mention `detect SSOTs`, got: %v", err)
	}
}
