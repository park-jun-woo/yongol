//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCheckSqlcOutPath — sqlc.yaml out 경로 검증의 read/parse/empty/mismatch/match 분기
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSqlc(t *testing.T, specsDir, content string) {
	t.Helper()
	dbDir := filepath.Join(specsDir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(content), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
}
