//ff:func feature=cli-init type=test control=sequence
//ff:what TestCopyFeaturesYAML — 정상 복사 / src 읽기 에러 / dest 쓰기(specs 부재) 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFeaturesYAML_Success(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	src := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(src, []byte("features: []\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyFeaturesYAML(target, src); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", "features.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "features: []\n" {
		t.Errorf("unexpected copied content: %q", got)
	}
}
