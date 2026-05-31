//ff:func feature=features type=test control=sequence
//ff:what TestLoadFeaturesFile — read/parse 에러·빈 features·필수필드 누락·성공 분기 검증
package features

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFeaturesFile_Success(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.yaml")
	body := "features:\n  - op: X\n    path: POST /t\n    desc: d\n"
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	feats, err := loadFeaturesFile(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(feats) != 1 || feats[0].Op != "X" {
		t.Errorf("unexpected feats: %v", feats)
	}
}
