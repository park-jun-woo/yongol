//ff:func feature=features type=test control=iteration dimension=1
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증

package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFeats(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

const baseFeats = `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
`

func TestRunAdd_NewFeaturesParseError(t *testing.T) {
	var out bytes.Buffer
	if err := RunAdd(&out, t.TempDir(), filepath.Join(t.TempDir(), "missing.yaml")); err == nil ||
		!strings.Contains(err.Error(), "new features") {
		t.Fatalf("want new features error, got %v", err)
	}
}

func TestRunAdd_ExistingParseError(t *testing.T) {
	dir := t.TempDir()
	newPath := filepath.Join(dir, "new.yaml")
	writeFeats(t, newPath, baseFeats)
	// specsDir has no features.yaml -> existing load fails.
	var out bytes.Buffer
	if err := RunAdd(&out, dir, newPath); err == nil ||
		!strings.Contains(err.Error(), "existing features") {
		t.Fatalf("want existing features error, got %v", err)
	}
}

func TestRunAdd_NoNewFeatures(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, baseFeats) // identical ops -> no diff
	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "no new features") {
		t.Errorf("want no-new-features message, got %q", out.String())
	}
}

func TestRunAdd_CreatesStubAndHash(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	newContent := baseFeats + `  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, newContent)

	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// SSaC stub created under service/task/DeleteTask.ssac
	stub := filepath.Join(specs, "service", "task", "DeleteTask.ssac")
	if _, err := os.Stat(stub); err != nil {
		t.Fatalf("expected stub at %s: %v", stub, err)
	}
	// .yongol hash written.
	if _, err := os.Stat(filepath.Join(specs, ".yongol")); err != nil {
		t.Errorf("expected .yongol hash: %v", err)
	}
	// features.yaml replaced with new content.
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if !strings.Contains(string(got), "DeleteTask") {
		t.Errorf("features.yaml not replaced: %q", got)
	}
	if !strings.Contains(out.String(), "1 new feature") {
		t.Errorf("want summary, got %q", out.String())
	}
}

func TestRunAdd_SkipExistingStub(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	// Pre-create the stub so the new op's write is skipped.
	stubDir := filepath.Join(specs, "service", "task")
	if err := os.MkdirAll(stubDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFeats(t, filepath.Join(stubDir, "DeleteTask.ssac"), "existing")
	newContent := baseFeats + `  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, newContent)

	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skip") {
		t.Errorf("want skip message, got %q", out.String())
	}
	// Pre-existing stub content preserved.
	got, _ := os.ReadFile(filepath.Join(stubDir, "DeleteTask.ssac"))
	if string(got) != "existing" {
		t.Errorf("stub overwritten: %q", got)
	}
}
