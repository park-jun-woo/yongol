//ff:func feature=contract type=test control=sequence
//ff:what TestDetectPreservedBranches — read 에러/주석 없음/hash 빈값/일치(Untouched) 분기 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestDetectPreserved_ReadError(t *testing.T) {
	if _, err := DetectPreserved(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error for missing file")
	}
}
