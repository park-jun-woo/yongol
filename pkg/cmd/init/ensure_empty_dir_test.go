//ff:func feature=cli-init type=test control=sequence
//ff:what TestEnsureEmptyDir — 부재/파일/force/빈디렉/비어있지않음 분기 검증

package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureEmptyDir_NotExist(t *testing.T) {
	if err := ensureEmptyDir(filepath.Join(t.TempDir(), "nope"), false); err != nil {
		t.Fatalf("missing dir should be ok: %v", err)
	}
}

func TestEnsureEmptyDir_NotADir(t *testing.T) {
	f := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureEmptyDir(f, false); err == nil {
		t.Fatal("want error when path is a file")
	}
}

func TestEnsureEmptyDir_EmptyOK(t *testing.T) {
	if err := ensureEmptyDir(t.TempDir(), false); err != nil {
		t.Fatalf("empty dir should be ok: %v", err)
	}
}

func TestEnsureEmptyDir_NonEmptyRejected(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "f"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureEmptyDir(dir, false); err == nil {
		t.Fatal("want error for non-empty dir without force")
	}
}

func TestEnsureEmptyDir_ForceOverridesNonEmpty(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "f"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureEmptyDir(dir, true); err != nil {
		t.Fatalf("force should override: %v", err)
	}
}
