//ff:func feature=migration type=test control=sequence
//ff:what TestGenerate_IncrementalMode — 스냅샷 존재 상태에서 DDL 변경 시 0002_* 파일 생성
package migration

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerate_IncrementalMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '');
`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '', age INTEGER NOT NULL DEFAULT 0);
`)
	res, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err != nil {
		t.Fatalf("incremental: %v", err)
	}
	if res.Mode != ModeIncremental {
		t.Errorf("expected incremental, got %v", res.Mode)
	}
	if !strings.Contains(res.MigrationFile, "0002_") {
		t.Errorf("expected 0002_* file, got %q", res.MigrationFile)
	}
}
