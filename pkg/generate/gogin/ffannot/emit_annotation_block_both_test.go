//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEmitAnnotationBlockBoth — Func+Type 동시 설정 및 빈 What 분기 검증
package ffannot

import (
	"strings"
	"testing"
)

func TestEmitAnnotationBlockBoth(t *testing.T) {
	got := EmitAnnotationBlock(Block{
		Func: FuncAnnot{Feature: "service", Type: "handler", Control: "sequence"},
		Type: TypeAnnot{Feature: "model", Type: "model"},
		What: "Combined — func and type",
	})
	if !strings.HasPrefix(got, "//ff:func feature=service type=handler control=sequence\n") {
		t.Fatalf("expected //ff:func first, got %q", got)
	}
	if !strings.Contains(got, "//ff:type feature=model type=model\n") {
		t.Fatalf("expected //ff:type line, got %q", got)
	}
	if !strings.Contains(got, "//ff:what Combined — func and type\n") {
		t.Fatalf("expected //ff:what line, got %q", got)
	}
}
