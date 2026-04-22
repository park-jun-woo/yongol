//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV10PanicNolint — `// nolint:panic` 주석으로 PRV-10 면제

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV10PanicNolint(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nfunc ActivateWorkflow() {\n"+
			"\t// nolint:panic\n"+
			"\tpanic(\"intentional\")\n}\n")
	diags := prv10PreservedPanic([]string{p})
	if len(diags) != 0 {
		t.Fatalf("nolint:panic should suppress, got %+v", diags)
	}
}
