//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: status 서브커맨드 end-to-end 2 케이스 (happy / missing-dir)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationStatus_Happy runs `yongol status <zenflow-specs>` and
// expects exit 0 with an SSOT Summary header. status is a read-only
// dashboard, so no arts-dir means no drift / artifacts sections.
func TestIntegrationStatus_Happy(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "status", specs)
	if err != nil {
		t.Fatalf("status happy: unexpected error: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.Contains(stdout, "SSOT Summary") {
		t.Errorf("expected stdout to contain `SSOT Summary`, got:\n%s", stdout)
	}
	// OpenAPI endpoints row is always rendered for a populated zenflow tree.
	if !strings.Contains(stdout, "OpenAPI") {
		t.Errorf("expected stdout to mention OpenAPI in summary, got:\n%s", stdout)
	}
}

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
