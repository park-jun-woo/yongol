//ff:func feature=contract type=test control=sequence
//ff:what test: TestStripAnnotationBlockRemovesHeader — //ff: 블록이 있는 파일에서 package 이하만 남는지 검증

package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlockRemovesHeader(t *testing.T) {
	in := []byte("//ff:func feature=x type=y control=sequence\n//ff:what demo\n//ff:checked llm=yongol-gen hash=abcdef01\npackage demo\n\nfunc Demo() {}\n")
	want := []byte("package demo\n\nfunc Demo() {}\n")
	got := StripAnnotationBlock(in)
	if !bytes.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
