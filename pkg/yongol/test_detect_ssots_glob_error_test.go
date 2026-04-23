//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — filepath.Glob ErrBadPattern 계약 회귀 (Phase012)
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
