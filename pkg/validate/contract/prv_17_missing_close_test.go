//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV17MissingClose — os.Open 후 defer Close 누락 시 ERROR

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV17MissingClose(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"os\"\n\n"+
			"func h() error {\n"+
			"\tf, err := os.Open(\"/tmp/x\")\n"+
			"\tif err != nil { return err }\n"+
			"\t_ = f\n"+
			"\treturn nil\n}\n")
	diags := prv17PreservedMissingClose([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
