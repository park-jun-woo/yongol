//ff:func feature=cli-featcheck type=test control=iteration dimension=1
//ff:what TestRunErrorBranches — read 실패 / parse 실패 / 빈 features / path·desc 누락 분기 검증

package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_ReadError(t *testing.T) {
	_, _, err := Run(filepath.Join(t.TempDir(), "does-not-exist.yaml"))
	if err == nil || !strings.Contains(err.Error(), "read features") {
		t.Fatalf("want read error, got %v", err)
	}
}

func TestRun_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(path, []byte("features: [unterminated"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "parse features") {
		t.Fatalf("want parse error, got %v", err)
	}
}

func TestRun_NoFeatures(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(path, []byte("tables: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "no features") {
		t.Fatalf("want no-features error, got %v", err)
	}
}

func TestRun_MissingPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	content := `features:
  - op: CreateTask
    path: ""
    desc: d
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "missing required field 'path'") {
		t.Fatalf("want missing path error, got %v", err)
	}
}

func TestRun_MissingDesc(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: ""
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "missing required field 'desc'") {
		t.Fatalf("want missing desc error, got %v", err)
	}
}
