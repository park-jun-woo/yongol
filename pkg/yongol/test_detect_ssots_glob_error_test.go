//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — filepath.Glob 오류 전파 계약 회귀 (Phase012)
package yongol

import (
	"errors"
	"path/filepath"
	"testing"
)

// TestDetectSSOTsGlobBadPatternContract pins the standard-library contract that
// Phase012's defensive guard relies on: filepath.Glob returns
// filepath.ErrBadPattern for syntactically invalid patterns. DetectSSOTs hard-
// codes its patterns so this path is effectively unreachable, but the guard
// exists to prevent silent pass if a future refactor introduces a caller-
// supplied pattern. If this contract ever changes the guard must be revisited.
func TestDetectSSOTsGlobBadPatternContract(t *testing.T) {
	// Unclosed character class — the canonical ErrBadPattern trigger.
	_, err := filepath.Glob("[")
	if err == nil {
		t.Fatalf("expected filepath.Glob(\"[\") to return ErrBadPattern, got nil")
	}
	if !errors.Is(err, filepath.ErrBadPattern) {
		t.Fatalf("expected filepath.ErrBadPattern, got %v", err)
	}
}

// TestDetectSSOTsHappyPathNoGlobError is a smoke regression — the hard-coded
// patterns in DetectSSOTs must never trigger filepath.ErrBadPattern across the
// matrix of detected directories. If any pattern in detect_ssots.go is changed
// to something syntactically invalid this test will fail at the DetectSSOTs
// call instead of silently returning empty matches.
func TestDetectSSOTsHappyPathNoGlobError(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	// Populate each SSOT directory so all glob patterns are exercised.
	writeFile(t, filepath.Join(tmp, "manifest.yaml"), minimalManifest)
	writeFile(t, filepath.Join(tmp, "api", "openapi.yaml"), "openapi: 3.0.0\n")
	writeFile(t, filepath.Join(tmp, "db", "schema.sql"), "-- noop\n")
	writeFile(t, filepath.Join(tmp, "service", "svc.ssac"), "// noop\n")
	writeFile(t, filepath.Join(tmp, "frontend", "index.html"), "<html></html>\n")
	writeFile(t, filepath.Join(tmp, "states", "s.md"), "# noop\n")
	writeFile(t, filepath.Join(tmp, "policy", "p.rego"), "package p\n")
	writeFile(t, filepath.Join(tmp, "tests", "scenario-a.hurl"), "GET /\n")
	writeFile(t, filepath.Join(tmp, "func", "pkg", "f.go"), "package pkg\n")

	if _, err := DetectSSOTs(tmp); err != nil {
		t.Fatalf("DetectSSOTs unexpectedly failed: %v", err)
	}
}
