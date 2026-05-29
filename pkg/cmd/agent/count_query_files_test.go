//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestCountQueryFiles — db/queries 하위 파일만 세고 디렉토리/부재는 0 반환 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountQueryFiles(t *testing.T) {
	// Missing directory returns 0.
	if got := countQueryFiles(t.TempDir()); got != 0 {
		t.Errorf("missing dir = %d, want 0", got)
	}

	specs := t.TempDir()
	qdir := filepath.Join(specs, "db", "queries")
	if err := os.MkdirAll(qdir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"a.sql", "b.sql"} {
		if err := os.WriteFile(filepath.Join(qdir, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// A subdirectory should not be counted.
	if err := os.MkdirAll(filepath.Join(qdir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if got := countQueryFiles(specs); got != 2 {
		t.Errorf("countQueryFiles = %d, want 2", got)
	}
}
