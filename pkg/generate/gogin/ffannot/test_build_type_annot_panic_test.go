//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestBuildTypeAnnotPanicsOnEmpty — 필수 필드 누락 시 panic

package ffannot

import "testing"

func TestBuildTypeAnnotPanicsOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic on empty type")
		}
	}()
	BuildTypeAnnot(TypeAnnot{Feature: "model"})
}
