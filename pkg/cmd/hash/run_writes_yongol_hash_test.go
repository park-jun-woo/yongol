//ff:func feature=cli-hash type=test control=sequence
//ff:what TestRun_WritesYongolHash — Run 실행 후 .yongol 파일에 sha256 해시 기록 확인

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
	featContent := []byte(`tables:
  tasks: {}
features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
    table: tasks
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
