//ff:func feature=contract type=test control=sequence
//ff:what TestExtractSignatureBranches — read/parse 에러·func 없음(zero)·init skip 분기 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestExtractSignature_ReadError(t *testing.T) {
	if _, err := ExtractSignature(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error")
	}
}
