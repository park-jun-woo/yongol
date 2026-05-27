//ff:func feature=cli type=test control=sequence
//ff:what latestSQLFileName test — .sql 파일 선택 검증

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLatestSQLFileName(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		got := latestSQLFileName(nil)
		if got != "" {
			t.Errorf("expected empty, got '%s'", got)
		}
	})
	t.Run("Multiple", func(t *testing.T) {
		dir := t.TempDir()
		for _, name := range []string{"001_init.sql", "002_add.sql", "003_drop.sql", "readme.md"} {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(""), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		got := latestSQLFileName(entries)
		if got != "003_drop.sql" {
			t.Errorf("expected '003_drop.sql', got '%s'", got)
		}
	})
	t.Run("NoSQL", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "readme.md"), []byte(""), 0o644); err != nil {
			t.Fatal(err)
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		got := latestSQLFileName(entries)
		if got != "" {
			t.Errorf("expected empty, got '%s'", got)
		}
	})
}
