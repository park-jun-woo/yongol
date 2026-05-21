//ff:func feature=cli-featcheck type=test control=sequence
//ff:what Run 테스트 — 정상/중복op/has_many참조오류/필수필드누락

package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_Happy(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: GetTask
    path: GET /tasks/{id}
    desc: Get a task
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	ff, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) != 0 {
		t.Errorf("want 0 diagnostics, got %d: %v", len(diags), diags)
	}
	if len(ff.Features) != 2 {
		t.Errorf("want 2 features, got %d", len(ff.Features))
	}
}

func TestRun_DuplicateOp_FT01(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: CreateTask
    path: POST /tasks/v2
    desc: Create a task v2
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) == 0 {
		t.Fatal("want diagnostics for duplicate op, got 0")
	}
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "[FT-01]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("want [FT-01] diagnostic, got %v", diags)
	}
}

func TestRun_HasManyRefError_FT10(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `tables:
  projects:
    has_many:
      - tasks
      - ghost_table
features:
  - op: CreateProject
    path: POST /projects
    desc: Create a project
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) == 0 {
		t.Fatal("want diagnostics for has_many ref error, got 0")
	}
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "[FT-10]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("want [FT-10] diagnostic, got %v", diags)
	}
}

func TestRun_MissingOp_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `features:
  - op: ""
    path: POST /tasks
    desc: Create a task
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, _, err := Run(path)
	if err == nil {
		t.Fatal("want error for missing op, got nil")
	}
	if !strings.Contains(err.Error(), "missing required field 'op'") {
		t.Errorf("unexpected error: %v", err)
	}
}
