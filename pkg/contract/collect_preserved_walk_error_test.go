//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectPreserved — 디렉토리 walk 시 preserved 파일만 반환
package contract

import (
	"path/filepath"
	"testing"
)

func TestCollectPreserved_WalkError(t *testing.T) {
	// Non-existent root -> WalkDir invokes fn with an error, which is returned.
	if _, err := CollectPreserved(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("want walk error for missing root")
	}
}
