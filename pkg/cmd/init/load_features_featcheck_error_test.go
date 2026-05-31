//ff:func feature=cli-init type=test control=sequence
//ff:what TestLoadFeatures — featcheck 에러 / ERROR 진단 / 정상 로드 분기 검증
package cliinit

import (
	"path/filepath"
	"testing"
)

func TestLoadFeatures_FeatcheckError(t *testing.T) {
	// Missing file -> featcheck.Run returns an error, propagated.
	if _, err := loadFeatures(filepath.Join(t.TempDir(), "nope.yaml")); err == nil {
		t.Fatal("want error for missing features.yaml")
	}
}
