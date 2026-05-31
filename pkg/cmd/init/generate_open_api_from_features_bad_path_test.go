//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateOpenAPIFromFeatures — path 그룹핑 stub 생성 / 잘못된 path 에러 / write 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateOpenAPIFromFeatures_BadPath(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs", "api"), 0o755); err != nil {
		t.Fatal(err)
	}
	feats := []features.Feature{{Op: "X", Path: "GARBAGE"}}
	if err := generateOpenAPIFromFeatures(target, templateData{}, feats); err == nil {
		t.Fatal("want error for invalid path")
	}
}
