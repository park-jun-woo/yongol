//ff:func feature=contract type=test control=sequence
//ff:what test: TestStripAnnotationBlockHandlesBanner — generator 배너와 //ff: 블록 모두 제거 확인

package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlockHandlesBanner(t *testing.T) {
	in := []byte("// Code generated DO NOT EDIT.\n\n//ff:type feature=api type=model\n//ff:what Foo defines model\npackage api\n\ntype Foo struct{}\n")
	want := []byte("package api\n\ntype Foo struct{}\n")
	got := StripAnnotationBlock(in)
	if !bytes.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
