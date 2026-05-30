//ff:func feature=features type=test control=iteration dimension=2
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증

package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const twoFeats = `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`

func setupSpecs(t *testing.T) string {
	t.Helper()
	specs := t.TempDir()
	if err := os.WriteFile(filepath.Join(specs, "features.yaml"), []byte(twoFeats), 0o644); err != nil {
		t.Fatal(err)
	}
	return specs
}

func TestRunRemove_NoOps(t *testing.T) {
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), t.TempDir(), nil, true); err == nil {
		t.Fatal("want error for empty ops")
	}
}

func TestRunRemove_ExistingLoadError(t *testing.T) {
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), t.TempDir(), []string{"X"}, true); err == nil ||
		!strings.Contains(err.Error(), "existing features") {
		t.Fatalf("want existing features error, got %v", err)
	}
}

func TestRunRemove_OpNotFound(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), specs, []string{"Ghost"}, true); err == nil ||
		!strings.Contains(err.Error(), "not found") {
		t.Fatalf("want not-found error, got %v", err)
	}
}

func TestRunRemove_Abort(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	// Answer "n" -> aborted, features.yaml untouched.
	if err := RunRemove(&out, strings.NewReader("n\n"), specs, []string{"DeleteTask"}, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "aborted") {
		t.Errorf("want aborted message, got %q", out.String())
	}
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if !strings.Contains(string(got), "DeleteTask") {
		t.Error("features.yaml should be untouched after abort")
	}
}

func TestRunRemove_ConfirmYesSkipsMissingSSaC(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	// Answer "y" -> proceeds; SSaC file doesn't exist -> skip branch.
	if err := RunRemove(&out, strings.NewReader("yes\n"), specs, []string{"DeleteTask"}, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skip") {
		t.Errorf("want skip message for missing ssac, got %q", out.String())
	}
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if strings.Contains(string(got), "DeleteTask") {
		t.Error("DeleteTask should be removed from features.yaml")
	}
	if !strings.Contains(string(got), "CreateTask") {
		t.Error("CreateTask should remain")
	}
}

func TestRunRemove_YesDeletesSSaC(t *testing.T) {
	specs := setupSpecs(t)
	// Create the SSaC file so the delete branch runs.
	ssacDir := filepath.Join(specs, "service", "task")
	if err := os.MkdirAll(ssacDir, 0o755); err != nil {
		t.Fatal(err)
	}
	ssacFile := filepath.Join(ssacDir, "DeleteTask.ssac")
	if err := os.WriteFile(ssacFile, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), specs, []string{"DeleteTask"}, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat(ssacFile); !os.IsNotExist(err) {
		t.Error("expected SSaC file to be deleted")
	}
	if !strings.Contains(out.String(), "1 feature(s) removed, 1 SSaC file(s) deleted") {
		t.Errorf("want summary, got %q", out.String())
	}
	// .yongol hash written.
	if _, err := os.Stat(filepath.Join(specs, ".yongol")); err != nil {
		t.Errorf("expected .yongol hash: %v", err)
	}
}
