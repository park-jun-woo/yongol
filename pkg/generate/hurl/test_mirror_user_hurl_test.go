//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorUserHurlFiles_CopiesScenarioAndInvariant — scenario-/invariant- 파일 복사 검증
package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorUserHurlFiles_CopiesScenarioAndInvariant asserts that
// scenario-*.hurl and invariant-*.hurl are mirrored verbatim into
// arts/tests/.
func TestMirrorUserHurlFiles_CopiesScenarioAndInvariant(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")

	testsSrc := filepath.Join(specsDir, "tests")
	if err := os.MkdirAll(testsSrc, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeFile(t, filepath.Join(testsSrc, "scenario-happy-path.hurl"), "# scenario body\n")
	writeFile(t, filepath.Join(testsSrc, "invariant-auth.hurl"), "# invariant body\n")
	// Should be ignored:
	writeFile(t, filepath.Join(testsSrc, "smoke.hurl"), "# stray user smoke\n")
	writeFile(t, filepath.Join(testsSrc, "README.md"), "noise\n")
	writeFile(t, filepath.Join(testsSrc, "random.hurl"), "# not a scenario/invariant prefix\n")

	if err := mirrorUserHurlFiles(specsDir, artsDir); err != nil {
		t.Fatalf("mirrorUserHurlFiles: %v", err)
	}

	mustExist(t, filepath.Join(artsDir, "tests", "scenario-happy-path.hurl"), "# scenario body\n")
	mustExist(t, filepath.Join(artsDir, "tests", "invariant-auth.hurl"), "# invariant body\n")
	mustNotExist(t, filepath.Join(artsDir, "tests", "smoke.hurl"))
	mustNotExist(t, filepath.Join(artsDir, "tests", "README.md"))
	mustNotExist(t, filepath.Join(artsDir, "tests", "random.hurl"))
}
