//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorSpecsTests_MissingTestsDirIsNoop — specs/tests/ 부재 시 no-op

package hurl_mirror

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorSpecsTests_MissingTestsDirIsNoop verifies that a project
// without specs/tests/ yields no error and no output — consistent with
// H-2 handling of the empty-tests case.
func TestMirrorSpecsTests_MissingTestsDirIsNoop(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")
	if err := os.MkdirAll(specsDir, 0o755); err != nil {
		t.Fatalf("mkdir specs: %v", err)
	}
	n, err := MirrorSpecsTests(specsDir, artsDir)
	if err != nil {
		t.Fatalf("MirrorSpecsTests: %v", err)
	}
	if n != 0 {
		t.Fatalf("mirrored = %d; want 0", n)
	}
	if _, err := os.Stat(filepath.Join(artsDir, "tests")); !os.IsNotExist(err) {
		t.Fatalf("arts/tests should not exist; stat err=%v", err)
	}
}
