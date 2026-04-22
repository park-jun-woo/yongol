//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: validate subcommand end-to-end 4 cases (happy / missing-dir / missing-args / parse-fail)

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationValidate_Happy runs `yongol validate <zenflow-specs>` and
// expects a clean report (0 errors, 0 warnings) with exit 0. Regression gate
// for the end-to-end validate pipeline — any upstream parser/rule that
// starts flagging zenflow breaks this.
func TestIntegrationValidate_Happy(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs)
	if err != nil {
		t.Fatalf("validate happy: unexpected error: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.Contains(stdout, "0 errors, 0 warnings") {
		t.Errorf("expected stdout to contain `0 errors, 0 warnings`, got:\n%s", stdout)
	}
}

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

// TestIntegrationValidate_MissingArgs verifies exit 2 (usage error) when no
// specs-dir is supplied. PhaseC01 mapped cobra.RangeArgs failures through
// usageArgs → *usageError; this test enforces that contract.
func TestIntegrationValidate_MissingArgs(t *testing.T) {
	_, _, err := runCmd(t, "validate")
	if err == nil {
		t.Fatal("expected usage error for missing arg, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}

// TestIntegrationValidate_ParseFail writes a malformed api/openapi.yaml
// into a tmpdir and expects validate to surface `parse failed` with exit 1.
// Confirms parse errors short-circuit before the Validate step.
func TestIntegrationValidate_ParseFail(t *testing.T) {
	tmp := t.TempDir()
	apiDir := filepath.Join(tmp, "api")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatalf("mkdir api: %v", err)
	}
	// Intentionally malformed YAML — a mapping value on the root key.
	bad := []byte("invalid: ::: yaml: :\n")
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), bad, 0644); err != nil {
		t.Fatalf("write bad openapi.yaml: %v", err)
	}
	stdout, _, err := runCmd(t, "validate", tmp)
	if err == nil {
		t.Fatalf("expected parse error, got nil\nstdout:\n%s", stdout)
	}
	if isUsageError(err) {
		t.Fatalf("parse failure must be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "parse failed") {
		t.Errorf("expected error to contain `parse failed`, got: %v", err)
	}
	if !strings.Contains(stdout, "Parse Errors") {
		t.Errorf("expected stdout to print `Parse Errors` banner, got:\n%s", stdout)
	}
}
