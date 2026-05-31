//ff:func feature=features type=test control=sequence
//ff:what TestLoadFeaturesFile — read/parse 에러·빈 features·필수필드 누락·성공 분기 검증
package features

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFeaturesFile_ReadError(t *testing.T) {
	if _, err := loadFeaturesFile(filepath.Join(t.TempDir(), "nope.yaml")); err == nil ||
		!strings.Contains(err.Error(), "read features") {
		t.Fatalf("want read error, got %v", err)
	}
}
