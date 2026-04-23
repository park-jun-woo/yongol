//ff:func feature=orchestrator type=test-helper control=sequence
//ff:what writeSQL — tmp 디렉토리에 sql 파일 기록 후 경로 반환

package sqlc

import (
	"os"
	"path/filepath"
	"testing"
)

// writeSQL writes body to <dir>/<name> and returns the full path. Used by
// ParseFile / ParseDir test suites to fabricate fixture files.
func writeSQL(t *testing.T, dir, name, body string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(body), 0644); err != nil {
		t.Fatalf("WriteFile(%s): %v", p, err)
	}
	return p
}
