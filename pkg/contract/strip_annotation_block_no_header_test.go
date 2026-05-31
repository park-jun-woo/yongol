//ff:func feature=contract type=test control=sequence
//ff:what TestStripAnnotationBlockBranches — 빈입력/주석없음/마지막줄 개행없는 주석 분기 검증
package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlock_NoHeader(t *testing.T) {
	in := []byte("package api\n\ntype Foo struct{}\n")
	if got := StripAnnotationBlock(in); !bytes.Equal(got, in) {
		t.Errorf("no-header input modified: got %q", got)
	}
}
