//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateSSaCFromFeatures — 도메인별 stub 생성 / mkdir 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateSSaCFromFeatures_MkdirError(t *testing.T) {
	// Make targetDir a file so MkdirAll under it fails.
	parent := t.TempDir()
	target := filepath.Join(parent, "asfile")
	if err := os.WriteFile(target, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	feats := []features.Feature{{Op: "X", Path: "GET /x"}}
	if err := generateSSaCFromFeatures(target, feats); err == nil {
		t.Fatal("want mkdir error when target is a file")
	}
}
