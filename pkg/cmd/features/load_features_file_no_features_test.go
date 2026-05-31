//ff:func feature=features type=test control=sequence
//ff:what TestLoadFeaturesFile — read/parse 에러·빈 features·필수필드 누락·성공 분기 검증
package features

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFeaturesFile_NoFeatures(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.yaml")
	if err := os.WriteFile(p, []byte("tables: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadFeaturesFile(p); err == nil || !strings.Contains(err.Error(), "no features") {
		t.Fatalf("want no-features error, got %v", err)
	}
}
