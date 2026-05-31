//ff:func feature=contract type=test control=sequence
//ff:what TestParsePreserveReasonBranches — read 에러 / reason 추출 성공 분기 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestParsePreserveReason_ReadError(t *testing.T) {
	if _, err := ParsePreserveReason(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error for missing file")
	}
}
