package clihash

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_WritesYongolHash(t *testing.T) {
	dir := t.TempDir()
	featContent := []byte(`features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
`)
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), featContent, 0o644); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := Run(&buf, dir); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	// .yongol must exist and contain sha256:
	data, err := os.ReadFile(filepath.Join(dir, ".yongol"))
	if err != nil {
		t.Fatalf("read .yongol: %v", err)
	}
	if !strings.Contains(string(data), "sha256:") {
		t.Errorf(".yongol content missing sha256: got %q", data)
	}

	// stdout must mention the destination
	if !strings.Contains(buf.String(), ".yongol") {
		t.Errorf("stdout missing .yongol mention: %q", buf.String())
	}
}

func TestRun_MissingFeaturesYaml(t *testing.T) {
	dir := t.TempDir()

	var buf bytes.Buffer
	err := Run(&buf, dir)
	if err == nil {
		t.Fatal("expected error for missing features.yaml, got nil")
	}
	if !strings.Contains(err.Error(), "features.yaml") {
		t.Errorf("error should mention features.yaml: %v", err)
	}
}

func TestRun_DuplicateOp_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: CreateTask
    path: POST /tasks/v2
    desc: Create a task v2
`
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	err := Run(&buf, dir)
	if err == nil {
		t.Fatal("expected error for duplicate op, got nil")
	}
	if !strings.Contains(err.Error(), "[FT-01]") {
		t.Errorf("error should mention [FT-01]: %v", err)
	}

	// .yongol must not be created.
	if _, statErr := os.Stat(filepath.Join(dir, ".yongol")); statErr == nil {
		t.Error(".yongol should not be created when validation fails")
	}
}
