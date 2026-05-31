//ff:func feature=contract type=test control=sequence
//ff:what TestStripAnnotationBlockBranches — 빈입력/주석없음/마지막줄 개행없는 주석 분기 검증
package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlock_BlankThenBody(t *testing.T) {
	in := []byte("\n\npackage x\n")
	want := []byte("package x\n")
	if got := StripAnnotationBlock(in); !bytes.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
