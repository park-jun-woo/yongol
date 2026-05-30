//ff:func feature=contract type=test control=sequence
//ff:what TestStripAnnotationBlockBranches — 빈입력/주석없음/마지막줄 개행없는 주석 분기 검증

package contract

import (
	"bytes"
	"testing"
)

func TestStripAnnotationBlock_Empty(t *testing.T) {
	if got := StripAnnotationBlock([]byte{}); len(got) != 0 {
		t.Errorf("empty input -> %q, want empty", got)
	}
}

func TestStripAnnotationBlock_NoHeader(t *testing.T) {
	in := []byte("package api\n\ntype Foo struct{}\n")
	if got := StripAnnotationBlock(in); !bytes.Equal(got, in) {
		t.Errorf("no-header input modified: got %q", got)
	}
}

func TestStripAnnotationBlock_AllComments(t *testing.T) {
	// Every line is a comment with no trailing newline on the last -> exercises
	// the nl<0 branch; result is empty (nothing after the header).
	in := []byte("//ff:type a\n//ff:what b")
	if got := StripAnnotationBlock(in); len(got) != 0 {
		t.Errorf("all-comments -> %q, want empty", got)
	}
}

func TestStripAnnotationBlock_BlankThenBody(t *testing.T) {
	in := []byte("\n\npackage x\n")
	want := []byte("package x\n")
	if got := StripAnnotationBlock(in); !bytes.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
