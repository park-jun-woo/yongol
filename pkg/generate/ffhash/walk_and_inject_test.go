//ff:func feature=gen-ffhash type=test control=sequence
//ff:what TestWalkAndInject — 디렉토리 순회 inject: 미존재/파일/skip/정상 경로 검증

package ffhash

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestWalkAndInject(t *testing.T) {
	annotated := func() []byte {
		return []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
			"/" + "/ff:what Demo — demo function\n" +
			"package demo\n\n" +
			"func Demo(x int) int {\n\treturn x + 1\n}\n")
	}

	t.Run("RootNotExist", func(t *testing.T) {
		if err := WalkAndInject(filepath.Join(t.TempDir(), "missing"), nil); err != nil {
			t.Errorf("expected nil for missing root, got: %v", err)
		}
	})

	t.Run("RootIsFile", func(t *testing.T) {
		dir := t.TempDir()
		fp := filepath.Join(dir, "file.go")
		if err := os.WriteFile(fp, annotated(), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := WalkAndInject(fp, nil); err != nil {
			t.Errorf("expected nil for file root, got: %v", err)
		}
		// File root short-circuits before any rewrite.
		got, _ := os.ReadFile(fp)
		if bytes.Contains(got, []byte("/"+"/ff:checked")) {
			t.Errorf("file root should not be rewritten, got:\n%s", got)
		}
	})

	t.Run("StatErrorNotNotExist", func(t *testing.T) {
		// A path whose parent is a regular file yields ENOTDIR (not IsNotExist).
		dir := t.TempDir()
		fp := filepath.Join(dir, "file")
		if err := os.WriteFile(fp, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := WalkAndInject(filepath.Join(fp, "sub"), nil)
		if err == nil || os.IsNotExist(err) {
			t.Errorf("expected non-NotExist stat error, got: %v", err)
		}
	})

	t.Run("WalkCallbackError", func(t *testing.T) {
		root := t.TempDir()
		bad := filepath.Join(root, "noaccess")
		if err := os.MkdirAll(bad, 0o000); err != nil {
			t.Fatalf("setup: %v", err)
		}
		defer os.Chmod(bad, 0o755)
		err := WalkAndInject(root, nil)
		if err == nil {
			t.Skip("walk did not surface a permission error (likely running as root)")
		}
	})

	t.Run("InjectsGoSkipsOthersAndSkipPrefix", func(t *testing.T) {
		root := t.TempDir()
		// .go file at top level -> injected.
		goPath := filepath.Join(root, "a.go")
		if err := os.WriteFile(goPath, annotated(), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// non-.go file -> skipped.
		txtPath := filepath.Join(root, "notes.txt")
		if err := os.WriteFile(txtPath, []byte("hello"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// .go file under a skipped directory -> skipped.
		vendorDir := filepath.Join(root, "vendor")
		if err := os.MkdirAll(vendorDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		vendorGo := filepath.Join(vendorDir, "v.go")
		if err := os.WriteFile(vendorGo, annotated(), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}

		if err := WalkAndInject(root, []string{"vendor"}); err != nil {
			t.Fatalf("WalkAndInject error: %v", err)
		}

		if got, _ := os.ReadFile(goPath); !bytes.Contains(got, []byte("/"+"/ff:checked")) {
			t.Errorf("expected a.go injected, got:\n%s", got)
		}
		if got, _ := os.ReadFile(txtPath); !bytes.Equal(got, []byte("hello")) {
			t.Errorf("txt file should be untouched, got: %s", got)
		}
		if got, _ := os.ReadFile(vendorGo); bytes.Contains(got, []byte("/"+"/ff:checked")) {
			t.Errorf("vendor go file should be skipped, got:\n%s", got)
		}
	})
}
