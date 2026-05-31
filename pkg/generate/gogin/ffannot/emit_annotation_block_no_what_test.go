//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEmitAnnotationBlockBoth — Func+Type 동시 설정 및 빈 What 분기 검증
package ffannot

import (
	"strings"
	"testing"
)

func TestEmitAnnotationBlockNoWhat(t *testing.T) {
	// Func set but What empty -> BuildWhat returns "" and the what line is skipped.
	got := EmitAnnotationBlock(Block{
		Func: FuncAnnot{Feature: "service", Type: "handler", Control: "sequence"},
		What: "",
	})
	if strings.Contains(got, "//ff:what") {
		t.Fatalf("empty What must not emit //ff:what line, got %q", got)
	}
	if !strings.HasPrefix(got, "//ff:func ") {
		t.Fatalf("expected //ff:func line, got %q", got)
	}
}
