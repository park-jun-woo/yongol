//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEnsureUnique — 충돌 발생 시 _2/_3 suffix 부여 검증

package fffile

import "testing"

func TestEnsureUnique(t *testing.T) {
	used := map[string]bool{}
	if got := EnsureUnique("convert_workflow.go", used); got != "convert_workflow.go" {
		t.Fatalf("first call: got %q want convert_workflow.go", got)
	}
	if got := EnsureUnique("convert_workflow.go", used); got != "convert_workflow_2.go" {
		t.Fatalf("second call: got %q want convert_workflow_2.go", got)
	}
	if got := EnsureUnique("convert_workflow.go", used); got != "convert_workflow_3.go" {
		t.Fatalf("third call: got %q want convert_workflow_3.go", got)
	}
	if got := EnsureUnique("", used); got != "" {
		t.Fatalf("empty candidate: got %q want empty", got)
	}
	// nil used map -> candidate returned unchanged (cannot persist).
	if got := EnsureUnique("file.go", nil); got != "file.go" {
		t.Fatalf("nil used map: got %q want file.go", got)
	}
}
