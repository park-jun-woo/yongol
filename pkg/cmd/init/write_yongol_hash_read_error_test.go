//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteYongolHash — 정상 해시 기록 / features 읽기 에러 / specs 부재 write 에러 분기 검증
package cliinit

import (
	"path/filepath"
	"testing"
)

func TestWriteYongolHash_ReadError(t *testing.T) {
	if err := writeYongolHash(t.TempDir(), filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("want read error for missing features.yaml")
	}
}
