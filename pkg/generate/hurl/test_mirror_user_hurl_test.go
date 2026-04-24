//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorUserHurlFiles — scenario-/invariant- 파일만 arts/tests/ 로 복사되고 smoke.hurl/관련없는 파일은 제외되는지 검증
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

// TestMirrorUserHurlFiles_MissingTestsDirIsNoop ensures projects
// without specs/tests/ still succeed without creating arts/tests/.
func TestMirrorUserHurlFiles_MissingTestsDirIsNoop(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")
	if err := os.MkdirAll(specsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := mirrorUserHurlFiles(specsDir, artsDir); err != nil {
		t.Fatalf("mirrorUserHurlFiles: %v", err)
	}
	if _, err := os.Stat(filepath.Join(artsDir, "tests")); !os.IsNotExist(err) {
		t.Fatalf("expected no arts/tests/ when specs/tests/ missing; stat err=%v", err)
	}
}

// TestMirrorUserHurlFiles_EmptySpecsDirNoop ensures empty specsDir
// argument is a safe no-op (defensive path for missing Fullstack).
func TestMirrorUserHurlFiles_EmptySpecsDirNoop(t *testing.T) {
	if err := mirrorUserHurlFiles("", t.TempDir()); err != nil {
		t.Fatalf("mirrorUserHurlFiles(\"\"): %v", err)
	}
}

// TestMirrorUserHurlFiles_OnlyUnmatchedFilesDoesNotCreateDir ensures
// tests/ dir with only ignored files (e.g. stray smoke.hurl) does not
// materialize an empty arts/tests/.
func TestMirrorUserHurlFiles_OnlyUnmatchedFilesDoesNotCreateDir(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")
	testsSrc := filepath.Join(specsDir, "tests")
	if err := os.MkdirAll(testsSrc, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeFile(t, filepath.Join(testsSrc, "smoke.hurl"), "# stray\n")
	writeFile(t, filepath.Join(testsSrc, "notes.txt"), "noise\n")

	if err := mirrorUserHurlFiles(specsDir, artsDir); err != nil {
		t.Fatalf("mirrorUserHurlFiles: %v", err)
	}
	if _, err := os.Stat(filepath.Join(artsDir, "tests")); !os.IsNotExist(err) {
		t.Fatalf("expected arts/tests/ NOT to be created; stat err=%v", err)
	}
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func mustExist(t *testing.T, path, wantBody string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != wantBody {
		t.Fatalf("%s body = %q; want %q", path, string(got), wantBody)
	}
}

func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected %s to not exist; err=%v", path, err)
	}
}
