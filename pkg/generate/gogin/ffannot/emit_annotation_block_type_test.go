//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEmitAnnotationBlockType — //ff:type 분기 출력 검증

package ffannot

import (
	"strings"
	"testing"
)

func TestEmitAnnotationBlockType(t *testing.T) {
	got := EmitAnnotationBlock(Block{
		Type: TypeAnnot{Feature: "model", Type: "model"},
		What: "CurrentUser — JWT claims",
	})
	if !strings.HasPrefix(got, "//ff:type feature=model type=model\n") {
		t.Fatalf("expected //ff:type prefix, got %q", got)
	}
}
