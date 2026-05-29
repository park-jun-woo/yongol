//ff:func feature=cli type=test control=sequence
//ff:what status missing-dir — 존재하지 않는 디렉토리 시 detect SSOTs 에러 (exit 1)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationStatus_MissingDir expects exit 1 (not a usage error) with
// a `detect SSOTs` message when the specs-dir does not exist.
func TestIntegrationStatus_MissingDir(t *testing.T) {
	_, _, err := runCmd(t, "status", "/does/not/exist/yongol-status-test")
	if err == nil {
		t.Fatal("expected error for missing specs-dir, got nil")
	}
	if isUsageError(err) {
		t.Fatalf("missing dir should be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "detect SSOTs") {
		t.Errorf("expected error to mention `detect SSOTs`, got: %v", err)
	}
}
