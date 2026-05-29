//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorSpecsTests_CopiesAllHurlFiles — 모든 .hurl 파일 (하위 포함) 복사 확인

package hurl_mirror

import (
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
