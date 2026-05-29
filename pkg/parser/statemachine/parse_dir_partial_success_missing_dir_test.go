//ff:func feature=statemachine type=test control=sequence topic=states
//ff:what ParseDir — 존재하지 않는 states/ 디렉토리는 SILENT-OK (진단 0, diagrams 비어있음)

package statemachine

import (
	"path/filepath"
	"testing"
)

// TestParseDirPartialSuccessMissingDir verifies SILENT-OK behaviour for a
// missing states/ directory (optional SSOT): diags nil, diagrams nil.
func TestParseDirPartialSuccessMissingDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "does-not-exist")

	diagrams, diags := ParseDir(dir)

	if len(diags) != 0 {
		t.Fatalf("expected 0 diags for missing dir (SILENT-OK), got %d: %v", len(diags), diags)
	}
	if len(diagrams) != 0 {
		t.Fatalf("expected 0 diagrams for missing dir, got %d", len(diagrams))
	}
}
