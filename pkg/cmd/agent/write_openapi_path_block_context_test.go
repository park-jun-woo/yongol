//ff:func feature=agent type=test control=sequence
//ff:what TestWriteOpenAPIPathBlockContext — opID의 path 블록 기록, 파일/op 부재 시 무기록 검증

package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteOpenAPIPathBlockContext(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	writeOpenAPIPathBlockContext(&b, dir, "ListUsers")
	out := b.String()
	if !strings.Contains(out, "OpenAPI path block (ListUsers):") || !strings.Contains(out, "/users") {
		t.Errorf("block → %q", out)
	}

	// Unknown op: nothing.
	var b2 strings.Builder
	writeOpenAPIPathBlockContext(&b2, dir, "Unknown")
	if b2.Len() != 0 {
		t.Errorf("unknown op wrote %q", b2.String())
	}

	// Missing file: nothing.
	var b3 strings.Builder
	writeOpenAPIPathBlockContext(&b3, t.TempDir(), "ListUsers")
	if b3.Len() != 0 {
		t.Errorf("missing file wrote %q", b3.String())
	}
}
