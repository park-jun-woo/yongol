//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteYongolHash — 정상 해시 기록 / features 읽기 에러 / specs 부재 write 에러 분기 검증

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteYongolHash_Success(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	feat := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(feat, []byte("features: []\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeYongolHash(target, feat); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", ".yongol"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "sha256:") {
		t.Errorf("unexpected .yongol: %q", got)
	}
}

func TestWriteYongolHash_ReadError(t *testing.T) {
	if err := writeYongolHash(t.TempDir(), filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("want read error for missing features.yaml")
	}
}

func TestWriteYongolHash_WriteError(t *testing.T) {
	feat := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(feat, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	// target has no specs/ subdir -> write fails.
	if err := writeYongolHash(t.TempDir(), feat); err == nil {
		t.Fatal("want write error when specs/ missing")
	}
}
