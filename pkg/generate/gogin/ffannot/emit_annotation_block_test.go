//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEmitAnnotationBlockFunc — //ff:func 분기 출력 검증

package ffannot

import "testing"

func TestEmitAnnotationBlockFunc(t *testing.T) {
	got := EmitAnnotationBlock(Block{
		Func: FuncAnnot{Feature: "service", Type: "handler", Control: "sequence"},
		What: "ActivateWorkflow — activate a workflow",
	})
	want := "//ff:func feature=service type=handler control=sequence\n" +
		"//ff:what ActivateWorkflow — activate a workflow\n"
	if got != want {
		t.Fatalf("EmitAnnotationBlock() = %q, want %q", got, want)
	}
}
