//ff:func feature=gen-ffhash type=test control=sequence
//ff:what TestRewriteOne — 파일 읽기/preserve/inject/no-diff 경로별 rewriteOne 검증

package ffhash

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestRewriteOne(t *testing.T) {
	t.Run("ReadError", func(t *testing.T) {
		err := rewriteOne(filepath.Join(t.TempDir(), "does-not-exist.go"))
		if err == nil {
			t.Errorf("expected read error for missing file, got nil")
		}
	})

	t.Run("PreservedFileUntouched", func(t *testing.T) {
		// A checked hash that cannot match the recomputed body hash makes
		// DetectPreserved return StatePreserved, so rewriteOne must skip it.
		dir := t.TempDir()
		path := filepath.Join(dir, "preserved.go")
		src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
			"/" + "/ff:what Demo — demo function\n" +
			"/" + "/ff:checked llm=yongol-gen hash=deadbeef\n" +
			"package demo\n\n" +
			"func Demo(x int) int {\n\treturn x + 1\n}\n")
		if err := os.WriteFile(path, src, 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// Sanity: this file is detected as preserved.
		state, derr := contract.DetectPreserved(path)
		if derr != nil || state != contract.StatePreserved {
			t.Fatalf("setup expected StatePreserved, got state=%v err=%v", state, derr)
		}
		if err := rewriteOne(path); err != nil {
			t.Fatalf("rewriteOne error: %v", err)
		}
		got, _ := os.ReadFile(path)
		if !bytes.Equal(got, src) {
			t.Errorf("preserved file was modified:\n%s", got)
		}
	})

	t.Run("InjectsCheckedLine", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "fresh.go")
		src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
			"/" + "/ff:what Demo — demo function\n" +
			"package demo\n\n" +
			"func Demo(x int) int {\n\treturn x + 1\n}\n")
		if err := os.WriteFile(path, src, 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := rewriteOne(path); err != nil {
			t.Fatalf("rewriteOne error: %v", err)
		}
		got, _ := os.ReadFile(path)
		if !bytes.Contains(got, []byte("/"+"/ff:checked llm=yongol-gen hash=")) {
			t.Errorf("expected injected checked line, got:\n%s", got)
		}
	})

	t.Run("NoDiffNoFuncFile", func(t *testing.T) {
		// No func decl -> InjectCheckedLine returns input unchanged; the file
		// has no //ff:checked line so DetectPreserved is NotApplicable, and
		// rewriteOne short-circuits on bytes.Equal (no write).
		dir := t.TempDir()
		path := filepath.Join(dir, "nofunc.go")
		src := []byte("package demo\n")
		if err := os.WriteFile(path, src, 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := rewriteOne(path); err != nil {
			t.Fatalf("rewriteOne error: %v", err)
		}
		got, _ := os.ReadFile(path)
		if !bytes.Equal(got, src) {
			t.Errorf("expected no-diff file unchanged, got:\n%s", got)
		}
	})
}
