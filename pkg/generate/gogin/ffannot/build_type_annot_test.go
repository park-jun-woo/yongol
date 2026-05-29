//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestBuildTypeAnnot — //ff:type 조립 검증

package ffannot

import "testing"

func TestBuildTypeAnnot(t *testing.T) {
	got := BuildTypeAnnot(TypeAnnot{Feature: "model", Type: "model"})
	want := "//ff:type feature=model type=model"
	if got != want {
		t.Fatalf("BuildTypeAnnot() = %q, want %q", got, want)
	}
}
