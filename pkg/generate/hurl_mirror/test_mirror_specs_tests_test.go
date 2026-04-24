//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorSpecsTests — specs/tests/ 전체 미러링 + prune 검증

package hurl_mirror

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorSpecsTests_CopiesAllHurlFiles verifies that every .hurl file
// under specs/tests/ — including nested sub-directories — is mirrored
// byte-for-byte into arts/tests/, and that non-hurl files are ignored.
func TestMirrorSpecsTests_CopiesAllHurlFiles(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")

	mustWrite(t, filepath.Join(specsDir, "tests", "smoke.hurl"), "# smoke body\n")
	mustWrite(t, filepath.Join(specsDir, "tests", "scenario-happy.hurl"), "# happy\n")
	mustWrite(t, filepath.Join(specsDir, "tests", "invariant-auth.hurl"), "# auth inv\n")
	mustWrite(t, filepath.Join(specsDir, "tests", "sub", "nested.hurl"), "# nested\n")
	mustWrite(t, filepath.Join(specsDir, "tests", "README.md"), "noise\n")

	n, err := MirrorSpecsTests(specsDir, artsDir)
	if err != nil {
		t.Fatalf("MirrorSpecsTests: %v", err)
	}
	if n != 4 {
		t.Fatalf("mirrored count = %d; want 4", n)
	}
	mustExist(t, filepath.Join(artsDir, "tests", "smoke.hurl"), "# smoke body\n")
	mustExist(t, filepath.Join(artsDir, "tests", "scenario-happy.hurl"), "# happy\n")
	mustExist(t, filepath.Join(artsDir, "tests", "invariant-auth.hurl"), "# auth inv\n")
	mustExist(t, filepath.Join(artsDir, "tests", "sub", "nested.hurl"), "# nested\n")
	mustNotExist(t, filepath.Join(artsDir, "tests", "README.md"))
}

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

// TestMirrorSpecsTests_EmptySpecsDir verifies that specsDir="" short
// circuits without touching the filesystem.
func TestMirrorSpecsTests_EmptySpecsDir(t *testing.T) {
	n, err := MirrorSpecsTests("", "")
	if err != nil {
		t.Fatalf("MirrorSpecsTests: %v", err)
	}
	if n != 0 {
		t.Fatalf("mirrored = %d; want 0", n)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func mustExist(t *testing.T, path, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != want {
		t.Fatalf("content of %s = %q; want %q", path, string(got), want)
	}
}

func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected %s to not exist; stat err=%v", path, err)
	}
}
