//ff:func feature=contract type=test control=sequence
//ff:what test: TestStripAnnotationBlockNoHeader — 어노테이션 없는 파일은 원본 그대로 반환

package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlockNoHeader(t *testing.T) {
	in := []byte("package demo\nfunc Demo() {}\n")
	got := StripAnnotationBlock(in)
	if !bytes.Equal(got, in) {
		t.Errorf("got %q want %q", got, in)
	}
}
