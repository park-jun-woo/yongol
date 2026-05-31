//ff:func feature=contract type=test control=sequence
//ff:what TestExtractExternalSymbolsBranches — read/parse 에러·func 없는 파일(zero) 분기 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbols_ReadError(t *testing.T) {
	if _, err := ExtractExternalSymbols(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error")
	}
}
