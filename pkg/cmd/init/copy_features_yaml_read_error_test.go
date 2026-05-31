//ff:func feature=cli-init type=test control=sequence
//ff:what TestCopyFeaturesYAML — 정상 복사 / src 읽기 에러 / dest 쓰기(specs 부재) 에러 분기 검증
package cliinit

import (
	"path/filepath"
	"testing"
)

func TestCopyFeaturesYAML_ReadError(t *testing.T) {
	if err := copyFeaturesYAML(t.TempDir(), filepath.Join(t.TempDir(), "nope.yaml")); err == nil {
		t.Fatal("want read error for missing src")
	}
}
