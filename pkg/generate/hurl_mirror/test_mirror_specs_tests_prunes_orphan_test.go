//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorSpecsTests_PrunesOrphan — 이전 실행 산출 .hurl 자동 삭제

package hurl_mirror

import (
	"path/filepath"
	"testing"
)

// TestMirrorSpecsTests_PrunesOrphan verifies that stale *.hurl files in
// arts/tests/ from a previous run are removed when the matching spec is
// deleted.
func TestMirrorSpecsTests_PrunesOrphan(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")

	mustWrite(t, filepath.Join(specsDir, "tests", "smoke.hurl"), "# keep\n")
	// Leftover from a previous run — no matching spec file.
	mustWrite(t, filepath.Join(artsDir, "tests", "stale.hurl"), "# stale\n")

	n, err := MirrorSpecsTests(specsDir, artsDir)
	if err != nil {
		t.Fatalf("MirrorSpecsTests: %v", err)
	}
	if n != 1 {
		t.Fatalf("mirrored = %d; want 1", n)
	}
	mustExist(t, filepath.Join(artsDir, "tests", "smoke.hurl"), "# keep\n")
	mustNotExist(t, filepath.Join(artsDir, "tests", "stale.hurl"))
}
