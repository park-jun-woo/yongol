//ff:func feature=agent type=test control=sequence
//ff:what TestWriteDDLContext — DDL 존재 시 컨텍스트 기록, 빈 테이블/부재 시 무기록 검증

package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteDDLContext(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "db"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "db", "users.sql"), []byte("CREATE TABLE users (id UUID);"), 0o644); err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	writeDDLContext(&b, dir, "users")
	out := b.String()
	if !strings.Contains(out, "DDL (users.sql):") || !strings.Contains(out, "CREATE TABLE users") {
		t.Errorf("output missing DDL context: %q", out)
	}

	// Empty table name: nothing written.
	var b2 strings.Builder
	writeDDLContext(&b2, dir, "")
	if b2.Len() != 0 {
		t.Errorf("empty table wrote %q, want nothing", b2.String())
	}

	// Missing file: nothing written.
	var b3 strings.Builder
	writeDDLContext(&b3, dir, "missing")
	if b3.Len() != 0 {
		t.Errorf("missing file wrote %q, want nothing", b3.String())
	}
}
