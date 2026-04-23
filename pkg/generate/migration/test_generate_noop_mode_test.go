//ff:func feature=migration type=test control=sequence
//ff:what TestGenerate_NoopMode — 변경 없을 때 noop 모드 반환 + 추가 마이그레이션 파일 없음
package migration

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerate_NoopMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY);
`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	res, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err != nil {
		t.Fatalf("noop: %v", err)
	}
	if res.Mode != ModeNoop {
		t.Errorf("expected noop, got %v", res.Mode)
	}
	entries, _ := os.ReadDir(filepath.Join(artsDir, "db", "migrations"))
	if len(entries) != 1 {
		t.Errorf("expected exactly 1 migration file after noop, got %d: %+v", len(entries), entries)
	}
}
