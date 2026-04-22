//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV17ClosePresent — defer Close 존재 시 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV17ClosePresent(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"os\"\n\n"+
			"func h() error {\n"+
			"\tf, err := os.Open(\"/tmp/x\")\n"+
			"\tif err != nil { return err }\n"+
			"\tdefer f.Close()\n"+
			"\t_ = f\n"+
			"\treturn nil\n}\n")
	diags := prv17PreservedMissingClose([]string{p})
	if len(diags) != 0 {
		t.Fatalf("defer Close should satisfy, got %+v", diags)
	}
}
