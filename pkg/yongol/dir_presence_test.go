//ff:func feature=orchestrator type=test control=sequence
//ff:what TestIsDir/dirPresence/isYongolRoot — 디렉토리 판별·presence 매핑·루트 판별 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestDirPresence(t *testing.T) {
	dir := t.TempDir()

	if got := dirPresence(dir, 3); got != SSOTPopulated {
		t.Errorf("dirPresence(existing, 3) = %v, want SSOTPopulated", got)
	}
	if got := dirPresence(dir, 0); got != SSOTDeclared {
		t.Errorf("dirPresence(existing, 0) = %v, want SSOTDeclared", got)
	}
	missing := filepath.Join(dir, "nope")
	if got := dirPresence(missing, 0); got != SSOTAbsent {
		t.Errorf("dirPresence(missing, 0) = %v, want SSOTAbsent", got)
	}
	// File count > 0 wins even when the directory does not exist.
	if got := dirPresence(missing, 1); got != SSOTPopulated {
		t.Errorf("dirPresence(missing, 1) = %v, want SSOTPopulated", got)
	}
}
